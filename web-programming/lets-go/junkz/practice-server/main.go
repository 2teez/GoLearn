package main

import (
	"fmt"
	"net/http"
)


func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World!\n")
}

func main() {
	server := http.Server {
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", hello)
	server.ListenAndServe()
}
