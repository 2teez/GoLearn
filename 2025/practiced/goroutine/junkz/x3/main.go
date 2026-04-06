package main

import (
	"fmt"
	"sync"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func main() {
	money := 100
	mu := sync.Mutex{}

	go func() {
		stingy(&money, &mu)
	}()

	go func() {
		spend(&money, &mu)
	}()
	time.Sleep(1 * time.Second)
	mu.Lock()
	fmt.Println("Money in the account: ", money)
	mu.Unlock()
	//
	// using letter count frequency table
	mu2 := sync.Mutex{}
	freq := make([]int, 26)
	for i := 1000; i < 1030; i++ {
		url := fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i)
		go countLetters(url, freq, &mu2)
	}
	time.Sleep(1 * time.Second)
	mu2.Lock()
	defer mu2.Unlock()
	for ind := range len(letters) {
		fmt.Printf("%c-%d ", letters[ind], freq[ind])
	}
}
