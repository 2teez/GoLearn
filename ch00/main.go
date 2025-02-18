package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	for i := 0; i < 5; i++ {
		PrintGreet()
		PrintCounter(i)
	}
}

func PrintCounter(number int) {
	fmt.Println(number)
}

func PrintGreet() {
	fmt.Println("Hello Go!")
}
