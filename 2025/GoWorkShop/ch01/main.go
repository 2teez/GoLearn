package main

import (
	"fmt"
	_ "sort" // used to check go libs
)

func main() {
	for val := range 101 {
		if val == 0 {
			continue
		}
		fmt.Println(val, fizzBuzz(val))
	}

	nums := generateNumbers(11, 10)
	fmt.Println(nums)
	//sort.Ints(nums)
	sortNumbers(swapper, nums)
	fmt.Println(nums)
}
