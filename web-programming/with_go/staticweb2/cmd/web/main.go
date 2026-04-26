package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// using file server
	file_server := http.FileServer(http.Dir("public"))
	mux.Handle("/public/", http.StripPrefix("/public/", file_server))

	server := &http.Server{
		Addr:    "127.0.0.1:4200",
		Handler: mux,
	}
	mux.HandleFunc("/", home)

	log.Println("Started Server on: http://127.0.0.1:4200")
	server.ListenAndServe()
}
