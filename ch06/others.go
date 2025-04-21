package main

import "fmt"

type number interface {
	~int | ~int64 | ~int32 | ~float32 | ~float64 | ~string
}

func add[T number](s T, o ...T) T {
	var result T = s
	for _, val := range o {
		result = result + val
	}
	return result
}

func reduce[T number](s T, o []T) T {
	var result T = s
	for _, val := range o {
		result += val
	}
	return result
}

func doubleStr(s string) string {
	return s + s
}

type integer interface {
	fmt.Stringer
	~int | ~float32 | ~float64
}

func double[T integer](n T) T {
	return n * n
}

type Printable interface {
	//~int | ~float64
	fmt.Stringer
}

type myInt int

func (mi myInt) String() string {
	return fmt.Sprintf("%d", mi)
}

type myFloat float64

func (mf myFloat) String() string {
	return fmt.Sprintf("%f", mf)
}

type myNumber[T Printable] struct {
	value T
}

func (mi myNumber[T]) String() string {
	return fmt.Sprintf("%v", mi.value)
}

func printIt[T Printable](t myNumber[T]) {
	fmt.Println(t)
}
