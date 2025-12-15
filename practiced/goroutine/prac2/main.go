package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const (
	example    = "https://example.com"
	golang     = "https://golang.org/"
	golang_doc = "https://golang.org/doc"
)

func main() {

	go responseSize(example)
	go responseSize(golang)
	go responseSize(golang_doc)
	time.Sleep(5 * time.Second)
}

func responseSize(mUrl string) {
	fmt.Println("Getting", mUrl)
	response, err := http.Get(mUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(string(body)))
}
