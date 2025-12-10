package main

import (
	"fmt"
	"log"
	"numbers"
)

func main() {
	fmt.Println(numbers.Add(4, 7))
	if n, err := numbers.Divide(4, 0); err != nil {
		log.Println(err)
	} else {
		fmt.Println(n)
	}
}
