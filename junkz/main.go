package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{1, 2, 1, 2, 5, 8}
	min, max := minMax(numbers)
	fmt.Println(findMissingNumbers(min, max, numbers))
}

func minMax(numbers []int) (int, int) {
	var min, max int
	if len(numbers) == 0 {
		return min, max
	}
	min, max = numbers[0], numbers[0]
	for _, val := range numbers {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max
}

func findMissingNumbers(min, max int, number []int) []int {
	var result []int
	number = removeDuplicates(number)
	for i := min; i < max; i++ {
		if slices.Contains(number, i) {
			continue
		}
		result = append(result, i)
	}
	return result
}

func removeDuplicates(numbers []int) []int {
	newNumber := map[int]struct{}{}
	for _, key := range numbers {
		newNumber[key] = struct{}{}
	}
	var keys []int
	for key := range newNumber {
		keys = append(keys, key)
	}
	return keys
}
