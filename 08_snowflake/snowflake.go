package main

import (
	"errors"
	"log"
	"sync"
	"time"
)

const (
	epoch            int64 = 1577808000000                 // 开始时间戳：2020-01-01 00:00:00
	workerIdBits     int64 = 5                             // 机器id所占的位数
	datacenterIdBits int64 = 5                             // 数据标识id所占的位数
	maxWorkerId      int64 = -1 ^ (-1 << workerIdBits)     // 支持的最大机器id，结果是31 (这个移位算法可以很快的计算出几位二进制数所能表示的最大十进制数)
	maxDatacenterId  int64 = -1 ^ (-1 << datacenterIdBits) // 支持的最大数据标识id，结果是31
	sequenceBits     int64 = 12                            // 序列在id中占的位数

	workerIdShift      = sequenceBits                                   // 机器ID向左移12位
	datacenterIdShift  = sequenceBits + workerIdBits                    // 数据标识id向左移17位(12+5)
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits // 时间截向左移22位(5+5+12)

	sequenceMask int64 = -1 ^ (-1 << sequenceBits) // 生成序列的掩码，这里为4095 (0b111111111111=0xfff=4095)
)

var worker *SnowflakeIdWorker

func init() {
	w, err := createWorker(0, 0)
	if err != nil {
		log.Fatal(err)
		return
	}
	worker = w

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

type SnowflakeIdWorker struct {
	mutex         sync.Mutex // 添加互斥锁 确保并发安全
	lastTimestamp int64      // 上次生成ID的时间截
	workerId      int64      // 工作机器ID(0~31)
	datacenterId  int64      //数据中心ID(0~31)
	sequence      int64      // 毫秒内序列(0~4095)
}

func main() {
	log.Println(Id())

	ch := make(chan int64)
	nu := 100000

	for i := 0; i < nu; i++ {
		go func() {
			ch <- Id()
		}()
	}
	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < nu; i++ {
		// check repeat
		id := <-ch
		_, ok := m[id]
		if ok {
			continue
		}

		m[id] = i
	}

	log.Printf("plan %d, actual %d", nu, len(m))
}

func Id() int64 {
	id, err := worker.nextId()
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return id
}

func createWorker(wId int64, dId int64) (*SnowflakeIdWorker, error) {
	if wId < 0 || wId > maxWorkerId {
		return nil, errors.New("worker ID excess of quantity")
	}
	if dId < 0 || dId > maxDatacenterId {
		return nil, errors.New("datacenter ID excess of quantity")
	}
	// 生成一个新节点
	return &SnowflakeIdWorker{
		lastTimestamp: 0,
		workerId:      wId,
		datacenterId:  dId,
		sequence:      0,
	}, nil
}

func (w *SnowflakeIdWorker) nextId() (int64, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	// 获取生成时的时间戳 毫秒
	now := time.Now().UnixNano() / 1e6
	//如果当前时间小于上一次ID生成的时间戳，说明系统时钟回退过这个时候应当抛出异常
	if now < w.lastTimestamp {
		return 0, errors.New("clock moved backwards")
	}
	if w.lastTimestamp == now {
		w.sequence = (w.sequence + 1) & sequenceMask
		if w.sequence == 0 {
			// 阻塞到下一个毫秒，直到获得新的时间戳
			for now <= w.lastTimestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 当前时间与工作节点上一次生成ID的时间不一致 则需要重置工作节点生成ID的序号
		w.sequence = 0
	}
	// 将机器上一次生成ID的时间更新为当前时间
	w.lastTimestamp = now
	id := int64((now-epoch)<<timestampLeftShift | w.datacenterId<<datacenterIdShift | (w.workerId << workerIdShift) | w.sequence)
	return id, nil
}
