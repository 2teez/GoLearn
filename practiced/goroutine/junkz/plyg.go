package main

import "fmt"

func main() {
	values := []int{0, 2, 4, 6, 8}
	ch := make(chan int, len(values))
	for _, v := range values {
		go func() {
			ch <- v * 2
		}()
	}
	for range len(values) {
		fmt.Println(<-ch)
	}
}
