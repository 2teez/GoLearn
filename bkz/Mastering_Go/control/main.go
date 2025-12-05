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

	for ind, value := range os.Args[1:] {
		fmt.Println(ind+1, value)
	}

}
