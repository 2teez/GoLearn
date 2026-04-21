package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the id from the parameter from the query string and try to
	// convert it to an integer using strconv.Atoi function. If the conversion
	// fails or the id is negative, return a 404 Not Found response.
	id := r.URL.Query().Get("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", idInt)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// if r.Method != "POST" {
	// use http constants like http.MethodPost
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		// instead of the two method above use
		// http.Error like so
		// And use http const http.StatusMethodNotAllowed instead of 405
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet...."))
}
