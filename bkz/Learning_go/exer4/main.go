package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := getRandomNumbers(0, 100, 1)
	printArrays(nums)
	//
	evenNums := getRandomNumbers(0, 10, 2)
	printArrays(evenNums)

	specialPrint(getRandomNumbers(0, 10, 1))
	specialPrint(getRandomNumbersNonRepeated(1, 20, 1))
	printArrays(getRandomNumbersNonRepeated(1, 20, 1))
}

func getRandomNumbers(start, stop, step int) (result []int) {
	for i := start; i < stop; i += step {
		result = append(result, rand.Intn(stop-start)+start)
	}
	return result
}

func getRandomNumbersNonRepeated(start, stop, _step int) (result []int) {
	size := stop - start
	var counter int
	intSet := map[int][]interface{}{}
	for {
		if counter == size {
			break
		}
		value := rand.Intn(size)
		if _, ok := intSet[value]; !ok {
			intSet[value] = []interface{}{}
			counter++
		}
	}
	result = []int{} // could use make(result, 0)
	for k := range intSet {
		result = append(result, k)
	}
	return
}

func printArrays(vals []int) {
	for _, val := range vals {
		fmt.Printf("%v, ", val)
	}
	fmt.Println()
}

func specialPrint(nums []int) {
	for _, val := range nums {
		switch {
		case val%6 == 0:
			fmt.Println(val, " -> Six")
		case val%2 == 0:
			fmt.Println(val, " -> Two")
		case val%3 == 0:
			fmt.Println(val, " -> Three")
		default:
			fmt.Println(val, " -> Never mind.")
		}
	}
}
