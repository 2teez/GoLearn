package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// for loop
	intVals := make([]int, 100)

	for i := 0; i < 100; i++ {
		intVals[i] = rand.Intn(100)
	}
	fmt.Println(intVals)
	for _, val := range intVals {
		switch {
		case val%6 == 0:
			fmt.Println(val, "Six")
		case val%3 == 0:
			fmt.Println("Three")
		case val%2 == 0:
			fmt.Println("Two")
		default:
			fmt.Println(val, "Never mind")
		}
	}
}
