package main

import (
	"fmt"
	"sync"
)

func Run2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan string)
	var wg2 sync.WaitGroup

	wg2.Add(2)

	go func() {
		defer wg2.Done()
		defer close(ch1)
		for range 10 {
			ch1 <- r.Intn(10) + 1
		}
	}()
	go func() {
		defer wg2.Done()
		defer close(ch2)
		for range 10 {
			ch2 <- r.Intn(10) + 1
		}
	}()
	go func() {
		defer close(ch3)
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				ch3 <- fmt.Sprintf("From channel 1: %v", v)
			case v, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				ch3 <- fmt.Sprintf("From channel 2: %v", v)
			}
		}

	}()

	for v := range ch3 {
		fmt.Println(v)
	}

	wg2.Wait()

}
