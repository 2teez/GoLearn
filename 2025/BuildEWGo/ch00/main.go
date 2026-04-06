package main

import (
	"fmt"
	"os"

	"github.com/practice/addition/addition"
)

func main() {
	args := os.Args
	fmt.Println(args) // print all the cli options
	fmt.Println(args[1:])
	if len(args) < 2 {
		fmt.Printf("Error: %s", "cli parameters must be more than two ")
		os.Exit(1)
	}

	fmt.Println(addition.Sum(0, args[1:]))
}
