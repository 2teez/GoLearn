package main

import (
	"fmt"
	"net/http"
)


type Handler struct{}

func (h *Handler) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World!")
}


func main() {
	handler := Handler{}
	server := http.Server {
		Addr: "127.0.0.1:8080",
		Handler: &handler,
	}

	server.ListenAndServe()
}
