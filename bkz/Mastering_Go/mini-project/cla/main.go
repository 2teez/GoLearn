package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need more than one arguments...")
		return
	}

	var min, max float64
	for _, val := range args[1:] {
		val, err := strconv.ParseFloat(val, 64)
		if err != nil {
			continue
		}
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	fmt.Println("Min: ", min, "Max: ", max)
}
