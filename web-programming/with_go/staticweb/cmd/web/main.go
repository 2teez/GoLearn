package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	file_server := http.FileServer(http.Dir("public"))
	mux.Handle("/public/", http.StripPrefix("/public/", file_server))

	mux.HandleFunc("/", home)

	server := &http.Server{
		Addr:    "127.0.0.1:4100",
		Handler: mux,
	}

	log.Println("Server is running on http://127.0.0.1:4100")

	server.ListenAndServe()
}
