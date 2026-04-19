package main

import (
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	server := &http.Server {
		Addr: "127.0.0.1:4000",
		Handler: mux,
	}

	log.Println("Starting server on :4000")
	server.ListenAndServe()
}
