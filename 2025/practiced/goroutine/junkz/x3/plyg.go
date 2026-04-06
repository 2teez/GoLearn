package main

import (
	"fmt"
	"sync"
)

func stingy(money *int, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	for range 1000000 {
		*money += 10
	}
	fmt.Println("Stingy Done")
}

func spend(money *int, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	for range 1000000 {
		*money -= 10
	}
	fmt.Println("Spend Done")
}
