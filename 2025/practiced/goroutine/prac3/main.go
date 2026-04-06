package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// creating a channel in golang
	myFirstChannel := make(chan string)
	go greetings(myFirstChannel)
	fmt.Println(<-myFirstChannel)
	//
	odd, even := make(chan int), make(chan int)
	go generateNumber(odd, 1, 10, 2)
	go generateNumber(even, 0, 10, 2)
	for {
		o, ok1 := <-odd
		e, ok2 := <-even
		if !ok1 || !ok2 {
			break
		}
		fmt.Println("Odd ->", o, " Even ->", e)
	}

	// reading the size of each websites using channel in golang
	ch := make(chan int)
	defer close(ch)
	websites := []string{"https://example.com", "https://golang.org/", "https://golang.org/doc"}
	for _, website := range websites {
		go getResponse(ch, website)
	}
	for _ = range len(websites) {
		fmt.Println(<-ch)
	}
}

func greetings(msg chan string) {
	msg <- "Hello, there...!"
	close(msg)
}

func generateNumber(num chan int, start, limit, step int) {
	defer close(num)
	for i := start; i <= limit; i += step {
		num <- i
	}
}

func getResponse(ch chan int, mUrl string) {
	fmt.Println("Getting, ", mUrl)
	resp, err := http.Get(mUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != err {
		log.Fatal(err)
	}
	ch <- len(body)
}
