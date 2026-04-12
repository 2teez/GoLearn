package main

import (
	"ch13/reader"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Println("./args[0] <filename>")
		os.Exit(1)
	}
	file, err := reader.File(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := reader.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
}
