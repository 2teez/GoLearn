package main

import (
	"fmt"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Printable interface {
	String() string
}

func main() {
	fmt.Println(double(34.5))
	fmt.Println(double(0.054))
	fmt.Println(double(3))

}

func double[I Integer](i I) I {
	return i + i
}
