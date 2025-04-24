package main

import (
	"fmt"
)

func main() {
	const x = 34 // untyped constant value
	const y int = 5

	var a float64 = x
	var b float64 = float64(y)
	var c byte = x

	fmt.Println(x, a, y, b, c)

	const (
		i int     = 10
		f float64 = 2.35
	)
	fmt.Println(i, f)

	var (
		_b     byte   = byte(^uint8(0))
		smallI int32  = int32(^uint32(0) >> 1)
		bigI   uint64 = uint64(^uint(0) >> 1)
	)
	fmt.Println(_b, smallI, bigI)
	fmt.Println(_b+1, smallI+1, bigI+1)
}
