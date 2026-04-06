package main

import (
	"flag"
	"fmt"
	"strings"
)

// using flag
var n = flag.Bool("n", false, "omitting the newline")
var sep = flag.String("sep", " ", "string separator")

func main() {
	//
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	//
	fmt.Println("Hello, world!")
	returned_value := local_pointer()
	fmt.Printf("local pointer %p => %d\n", returned_value, *returned_value)
}

func local_pointer() *int {
	local := 23
	return &local
}
