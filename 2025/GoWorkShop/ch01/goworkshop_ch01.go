package main

import (
	"math/rand"
	"strconv"
)

type (
	ptFuncTy func(*int, *int)
	intSlice []int
)

func swapper(a, b *int) {
	if *a > *b {
		*b, *a = *a, *b
	}
}

func sortNumbers(f ptFuncTy, nums intSlice) {
	for range len(nums) - 1 {
		for i := range len(nums) - 1 {
			f(&nums[i], &nums[i+1])
		}
	}
}

func generateNumbers(times, ranged int) intSlice {
	values := make(intSlice, 0, times)
	for range times {
		values = append(values, rand.Intn(ranged)+1)
	}
	return values
}

func fizzBuzz(number int) string {
	switch {
	case number%15 == 0:
		return "FizzBuzz"
	case number%3 == 0:
		return "Fuzz"
	case number%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(number)
	}
}
