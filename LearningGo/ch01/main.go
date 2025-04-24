package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Int())
	first := 100
	second := first
	first++
	fmt.Println("First Value: ", first)
	fmt.Println("Second Value: ", second)
	///
	var pSecond *int = &first
	fmt.Println("First Value: ", first)
	fmt.Println("Second Value: ", pSecond, *pSecond)
}
