package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	LoadClient()
	concurrent := GlobalConfig.Http.Concurrent
	loop := GlobalConfig.Http.Loop
	limit := GlobalConfig.Http.Limit
	//wait := GlobalConfig.Http.Wait

	responseChan := make(chan []byte)
	defer close(responseChan)
	startTime := time.Now()

	for i := 0; i < concurrent; i++ {
		go func() {
			if limit > 0 {
				// 用令牌桶限制并发数量
				l := 0
				limiter := rate.NewLimiter(rate.Limit(limit), limit)
				limiter.AllowN(time.Now(), limit)
				for {
					if limiter.Allow() {
						got := Get("http://127.0.0.1:8080")
						responseChan <- got
						l++
					} else {
						time.Sleep(200 * time.Millisecond)
					}
					if l >= loop {
						break
					}
				}
			} else {
				// 并发无限制
				for j := 0; j < loop; j++ {
					got := Get("http://127.0.0.1:8080")
					responseChan <- got
				}
			}
		}()
	}

	counter := 0
	for {
		//time.Sleep(time.Duration(wait) * time.Second)
		select {
		case <-responseChan:
			counter++
			//TODO 这里插入查询结果验证逻辑
			if counter >= concurrent*loop {
				totalTime := time.Now().Sub(startTime).Nanoseconds()
				fmt.Printf("QPS %f", float32(concurrent*loop*1e9)/float32(totalTime))
				return
			}
		}
	}
}
