package main

import (
	"fmt"
	"github.com/hashicorp/mdns"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request for hello!")
	fmt.Fprint(w, "Hello from service!")
}

func startServer() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func main() {
	go startServer()

	for {
	}
}
