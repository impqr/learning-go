package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const (
	fileName = "config.yml"
)

var GlobalConfig Config

// Config 第一步：定义与配置文件匹配的结构体嵌套层级（理论上yml有多少级就需要多少个结构体）
type Config struct {
	Server ServerContent `yaml:"server"`
	Http   HttpContent   `yaml:"http"`
}

type ServerContent struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

type HttpContent struct {
	ReadTimeout         string `yaml:"readTimeout"`
	WriteTimeout        string `yaml:"writeTimeout"`
	MaxIdleConnDuration string `yaml:"maxIdleConnDuration"`
	Concurrent          int    `yaml:"concurrent"`
	Limit               int    `yaml:"limit"`
	Loop                int    `yaml:"loop"`
	Wait                int    `yaml:"wait"`
}

// 第二步：将读取的yml文件序列化为对应的结构体
func init() {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Panicln(err)
	}
	err = yaml.Unmarshal(bytes, &GlobalConfig)
	if err != nil {
		log.Panicln(err)
	}
}
