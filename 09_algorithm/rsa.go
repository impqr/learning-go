package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		_ = fmt.Errorf("error: %s", err)
	}

	pub := key.N
	pri := key.D

	fmt.Printf("rsa key pairs length: %d(pub) %d(pri)\n", pub.BitLen(), pri.BitLen())
	//fmt.Printf("rsa generate key pair: \npublic key: %s \nprivate key: %s\n", pub, pri)

	str := "Hello world! I'm M.C.Jackson!"

	bytes, err := rsa.EncryptPKCS1v15(rand.Reader, &(key.PublicKey), []byte(str))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("secret message length: ", len(bytes))

	origin, err := rsa.DecryptPKCS1v15(rand.Reader, key, bytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("origin string is: ", string(origin))
}
