package main

import (
	"fmt"
	"math/rand"

	"github.com/practice/arith/arith"
)

func main() {
	intArr := []int{4, 7, 2, 9}
	for _, val := range intArr {
		fmt.Printf("%p\n", &val)
	}

	if result, err := arith.Div(23, 0); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(doMaths(4, 9, arith.Add))

	for i := 1; i <= 10; i++ {
		a, b := rand.Intn(5), rand.Intn(10)
		fmt.Println(a, "divided by", b, doMaths(a, b, arith.Div))
	}
}

func doMaths[T arith.Integer](num1, num2 T,
	op arith.Arithmetics[T]) T {
	var result T
	if val, err := op(num1, num2); err != nil {
		defer func() {
			fmt.Printf("Error: %v\n", err)
		}()
	} else {
		result = val
	}
	return result
}
