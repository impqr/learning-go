package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	paths = []string{
		"/Users/pangqirong/Desktop/logger.log",
	}

	dbInit = map[string]string{
		"host":    "localhost",
		"port":    "3306",
		"user":    "root",
		"pass":    "123456",
		"scheme":  "test",
		"charset": "utf8",
	}
)

func main() {
	db := Connect(dbInit)
	defer db.Close()
	for _, path := range paths {
		read(path)
	}
}

func read(path string) {
	fmt.Println(path)

	file, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	for {
		bytes, b, err := buf.ReadLine()
		if err != nil {
			log.Panicln(err)
		}

		line := string(bytes)
		data := strings.Split(line, " ")

		if len(data) != 13 {
			continue
		}

		dateString := data[0] + " " + strings.ReplaceAll(data[1], ",", ".")
		t, err := time.ParseInLocation("2006-01-02 15:04:05.000", dateString, time.Local)
		if err != nil {
			log.Panicln(err)
		}
		timeLong := t.UnixNano() / 1e6

		action := data[9]
		id := data[10]
		attach := data[11:]

		fmt.Println(timeLong, action, id, attach)

		fmt.Println(b)

		break // only one line
	}
}
