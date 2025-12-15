package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	websites := map[string]string{
		"example":    "https://example.com",
		"golang":     "https://golang.org/",
		"golang-doc": "https://golang.org/doc",
	}
	for named := range websites {
		err := responseSize(websites[named])
		if err != nil {
			log.Fatal(err)
		}
	}
}

func responseSize(mUrl string) error {
	fmt.Println("Getting", mUrl)
	response, err := http.Get(mUrl)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	fmt.Println(len(string(body)))
	return nil
}
