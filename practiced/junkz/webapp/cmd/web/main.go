package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", home)

	log.Println("Starting server on :8118")
	err := http.ListenAndServe(":8118", server)
	log.Fatal(err)
}
