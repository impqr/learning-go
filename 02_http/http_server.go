package main

import (
	"fmt"
	"net/http"
)

func main() {
	addr := "127.0.0.1:8080"

	http.HandleFunc("/", mainHandler)
	fmt.Println("Server start at", addr)
	_ = http.ListenAndServe(addr, nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "you are request for ", r.RequestURI)
}
