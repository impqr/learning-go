package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	person := Person{Name: "张三", Age: 18}

	// 序列化
	data, err := json.Marshal(person)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(string(data))

	// 反序列化
	p := Person{}
	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(p)
}

// Person 结构体中的json标签用来映射序列化时显示的名称（反射），omitempty表示值不定义时省略输出（但会有默认值）
type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
