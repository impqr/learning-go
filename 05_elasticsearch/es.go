package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
)

func main() {
	es, _ := elasticsearch.NewClient(config)
	ping, _ := es.Ping()
	defer ping.Body.Close()
	log.Println(ping)
}

var config = elasticsearch.Config{
	Addresses: []string{"http://127.0.0.1:9200"},
	Username:  "root",
	Password:  "pass",
}
