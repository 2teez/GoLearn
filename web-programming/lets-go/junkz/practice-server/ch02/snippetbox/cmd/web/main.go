package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	server := &http.Server{
		Addr:    "127.0.0.1:4000",
		Handler: mux,
	}

	log.Println("Starting server on :4000")
	server.ListenAndServe()
}
