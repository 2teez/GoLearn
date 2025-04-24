package main

import (
	"errors"
	"fmt"
)

type parameters struct {
	start, stop, step int
}

type funcType func(int, int) (int, error)

var (
	add = func(a, b int) (int, error) { return a + b, nil }
	sub = func(a, b int) (int, error) { return a - b, nil }
	mul = func(a, b int) (int, error) { return a * b, nil }
	div = func(a, b int) (int, error) {
		if b == 0 {
			return 0, errors.New("Division by Zero")
		}
		return a + b, nil
	}
	mod = func(a, b int) int { return a % b }
)

func main() {
	fmt.Println("Start Here!")
	values := make([]int, 0)
	fmt.Println(values, cap(values), len(values))
	for i := 0; i < 10; i += 1 {
		values = append(values, i)
	}
	fmt.Println(values)
	fmt.Println(addTo(struct{ start, stop, step int }{0, 11, 1}))
	fmt.Println(addTo(parameters{step: 1, start: 0, stop: 101}))
	fmt.Println(addToSlice(3))
	fmt.Println(addToSlice(3, 2, 1, 8))
	fmt.Println(addToSlice(3, []int{5, 9, 0, 1}...))
	///
	printFunc([][]int{{3, 8}, {5, 9}, {2, 0}}, []funcType{add, sub, div}...)
	/// using prefixer
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}

func prefixer(s string) func(string) string {
	return func(actualString string) string {
		return s + " " + actualString
	}
}

func printFunc(values [][]int, f ...funcType) {
	for i, val := range values {
		res, err := f[i](val[0], val[1])
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	}
}

func addToSlice(start int, others ...int) []int {
	values := make([]int, 1, len(others)+1)
	values[0] = start
	for _, val := range others {
		values = append(values, val)
	}
	return values
}

func addTo(vals struct {
	start, stop, step int
}) int {
	var result int
	for i := vals.start; i < vals.stop; i += vals.step {
		result += i
	}
	return result
}
