package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	filehandler, _, err := r.FormFile("uploaded")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer filehandler.Close()

	data := make([]byte, 1024)
	n, err := filehandler.Read(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Uploaded file: %s", string(data[:n]))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	// handler function
	http.HandleFunc("/process", process)

	server.ListenAndServe()
}
