package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Microsecond)
	fmt.Printf("\r%d\n", Fibo(45))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func Fibo(num int) int {
	if num <= 2 {
		return num
	}
	return Fibo(num-2) + Fibo(num-1)
}
