package main

import (
	"fmt"
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
