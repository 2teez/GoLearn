package main

import (
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, World!\n"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
