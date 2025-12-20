package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

func write(w http.ResponseWriter, msg string) {
	_, err := fmt.Fprintf(w, "%s", msg)
	check(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	write(w, "Hello World")
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//write(w, "view signatures...")
	html, err := template.ParseFiles("./ui/html/pages/view.html")
	check(err)
	err = html.Execute(w, nil)
	check(err)
}
