package main

import (
	"log"
	"net/http"
)

func fs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello, fs..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", fs)

	file_server := http.FileServer(http.Dir("/public/"))
	mux.Handle("/public/", http.StripPrefix("/public", file_server))

	server := http.Server{
		Addr:    "127.0.0.1:4100",
		Handler: mux,
	}

	log.Println("Server is running on http://127.0.0.1:4100")

	server.ListenAndServe()
}
