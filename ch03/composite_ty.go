package main

import (
	"fmt"
)

func main() {
	var arr [3]string = [3]string{"java", "javascript", "clojure"}
	fmt.Println(arr, cap(arr), len(arr))
	for index := range arr {
		fmt.Println(index, arr[index])
	}
	arr1 := [3]int{1, 2, 4}
	fmt.Println(arr1)
	// slices
	arr2 := []string{}
	fmt.Println(arr2, len(arr2), cap(arr2))
	arr2 = append(arr2, []string{"clojure", "erlang", "crsytal"}...)
	fmt.Println(arr2)
	///
	nums := make([]int, 5)
	fmt.Println(nums)
	nums = append(nums, []int{6, 9, 0, 1}...)
	fmt.Println(nums)
	clear(nums)
	nums = []int{}
	fmt.Println(nums)
	nums = append(nums, []int{8, 3, 1}...)
	fmt.Println(nums)
	n := make([]int, 2)
	copy(n, nums[1:])
	fmt.Println(n, nums)
}
