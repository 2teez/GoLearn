package main

import (
	"fmt"
)

func main() {
	for val := range 101 {
		if val == 0 {
			continue
		}
		fmt.Println(val, fizzBuzz(val))
	}
}
