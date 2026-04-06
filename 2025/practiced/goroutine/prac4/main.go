package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type WebPage struct {
	mUrl string
	size int
}

func (wp WebPage) String() string {
	return fmt.Sprintf("%s: %d", wp.mUrl, wp.size)
}

func main() {
	websites := []string{
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
	}

	webChan := make(chan WebPage)
	defer close(webChan)
	for _, website := range websites {
		go getWebPageSize(webChan, website)
	}

	for _ = range len(websites) {
		fmt.Println(<-webChan)
	}
}

func getWebPageSize(webChan chan WebPage, webpage string) {
	resp, err := http.Get(webpage)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	webChan <- WebPage{size: len(body), mUrl: webpage}

}
