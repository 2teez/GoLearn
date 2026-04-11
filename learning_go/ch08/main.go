package main

import (
	"fmt"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type MyInt int
type MyFloat float64

func main() {
	fmt.Println(double(34.5))
	fmt.Println(double(0.054))
	fmt.Println(double(3))
	//
	var intValue MyInt = 3
	printIt(intValue)
	printIt(MyFloat(3.14))

}

func double[I Integer](i I) I {
	return i + i
}

func printIt[P Printable](p P) {
	fmt.Println(p)
}

func (i MyInt) String() string {
	return fmt.Sprintf("%v", int(i))
}

func (f MyFloat) String() string {
	return fmt.Sprintf("%v", float64(f))
}
