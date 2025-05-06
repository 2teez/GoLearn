package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting server on http://localhost:4220")
	if err := http.ListenAndServe(":4220", mux); err != nil {
		log.Fatalln(err)
	}
}
