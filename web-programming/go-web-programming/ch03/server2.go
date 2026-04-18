package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "Form data: %v", r.Form)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// handler function
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
