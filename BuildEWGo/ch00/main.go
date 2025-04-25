package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	fmt.Println(args) // print all the cli options
	fmt.Println(args[1:])
	if len(args) < 2 {
		fmt.Printf("Error: %s", "cli parameters must be more than two ")
		os.Exit(1)
	}

	fmt.Println(Sum(0, args[1:]))
}

func Sum(start int, o []string) int {
	result := start
	for _, value := range o {
		value, err := strconv.Atoi(value)
		if err != nil {
			defer func() {
				fmt.Println(fmt.Errorf("%v", err))
			}()
		}
		result += value
	}
	return result
}
