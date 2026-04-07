package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for loop
	intVals := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		intVals[i] = rand.Intn(100)
	}
	fmt.Println(intVals)
}
