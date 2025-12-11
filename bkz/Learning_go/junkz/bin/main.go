package main

import (
	"fmt"
	"junkz/pkg/src/readers"
	"log"
)

func main() {
	fmt.Println("Hello...")
	n, err := readers.ReadLine()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(n)
}
