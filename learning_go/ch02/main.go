package main

import (
	"fmt"
	"math"
)

func main() {
	i := 20
	f := float64(i)
	fmt.Printf("%d, %f\n", i, f)

	// untyped constant value assigable to
	// both integer and float values
	const value = 23
	var integerValue int = value
	var floatValue float64 = value
	fmt.Printf("%d, %.3f\n", integerValue, floatValue)
	// 3
	var (
		b      byte    = math.MaxUint8
		smallI int32   = math.MaxInt32
		bigI   float64 = math.MaxFloat64
	)
	fmt.Printf("\n%d, \n%d, \n%.3f\n", b+1, smallI+1, bigI+1)
}
