package main

import (
	"fmt"
	"slices"
)

func main() {
	numbers := []int{1, 2, 1, 2, 5, 8}
	min, max := min_max(numbers)
	fmt.Println(find_missing_number(min, max, numbers))
}

func min_max(numbers []int) (int, int) {
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

func find_missing_number(min, max int, number []int) []int {
	var result []int
	number = removeDuplicate(number)
	for i := min; i < max; i++ {
		if slices.Contains(number, i) {
			continue
		}
		result = append(result, i)
	}
	return result
}

func removeDuplicate(numbers []int) []int {
	newNumber := map[int]struct{}{}
	for _, key := range numbers {
		newNumber[key] = struct{}{}
	}
	var keys []int
	for key, _ := range newNumber {
		keys = append(keys, key)
	}
	return keys
}
