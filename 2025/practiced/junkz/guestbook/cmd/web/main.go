package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", home)
	server.HandleFunc("/home", home)
	server.HandleFunc("/guestbook/view", viewHandler)

	log.Println("Starting the server on http://localhost:6110")
	err := http.ListenAndServe(":6110", server)
	log.Fatal(err)
}
