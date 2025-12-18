package main

import (
	"fmt"
	"sync"
	"time"
)

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
	defer mu.Unlock()
	fmt.Println("Money in the account: ", money)
}
