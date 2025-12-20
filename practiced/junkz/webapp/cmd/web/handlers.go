package main

import (
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func write(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(message))
	check(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, Web...")
	write(w, "Hello, Web..")
}
