package main

import "fmt"

func main() {
	fmt.Println("Assignment in Golang...")
	fmt.Println(gcd(23, 3))
	fmt.Println(fibonacci(12))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fibonacci(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
