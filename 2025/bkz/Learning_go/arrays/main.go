package main

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

func main() {
	nums := []int{5, 8, 0, 1, 5}
	printArray(nums)

	x := []int{5, 9, 0}
	y := []int{5, 9, 0}
	z := []int{3, 0, 2, 4}
	//str := []string{"java", "javascript"}
	fmt.Println(slices.Equal(x, y))
	fmt.Println(slices.Equal(x, z))
	//fmt.Println(slices.Equal(str, y))
	fmt.Println(reflect.DeepEqual(z, y))
	//
	m := make([]int, 0)
	msgs := []string{"using the printAny function", "capacity of array: "}
	printAny(cap(m), msgs...)
	m = append(m, 2, 3)
	m = append(m, z...)
	printArray(m)
	printAny(cap(m), msgs...)
	printAny(cap(msgs))
	// using an array and slice
	primes := [5]int{2, 3, 5, 7, 11}
	printArray(primes[:]) // converting an array to slice..

}

func printArray(t []int) {
	for _, val := range t {
		fmt.Println(val)
	}
}

func printAny[T any](i T, msg ...string) {
	if len(msg) == 0 {
		fmt.Printf("%v \n", i)
		return
	}
	fmt.Printf("%s %v \n", strings.Join(msg, ", "), i)
}
