package anons

import (
	"fmt"
	"sync"
)

func RunNonAnons() {
	// anonymous function goroutine
	// without parameters
	var wg sync.WaitGroup
	wg.Add(3)
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func RunAnons() {
	// anonymous function goroutine
	// with parameters
	var wg sync.WaitGroup
	wg.Add(1)
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
