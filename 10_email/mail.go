package main

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

func main() {
	Send()
}

func Send() {
	m := gomail.NewMessage()
	m.SetHeader("From", "impqr@qq.com")
	m.SetHeader("To", "qirong.pang@icloud.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 465, "impqr@qq.com", "123456")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
