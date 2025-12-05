package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println(len(os.Args))
	if len(os.Args) != 3 {
		ctrl_name := strings.Split(os.Args[0], "/")
		fmt.Fprintf(os.Stderr, "./%s <option> <Arguments>\n", ctrl_name[len(ctrl_name)-1])
		return
	}

	// using old traditional for loop
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(i+1, os.Args[i])
	}
	// using forloop with range
	for ind, value := range os.Args[1:] {
		fmt.Println(ind+1, value)
	}

}
