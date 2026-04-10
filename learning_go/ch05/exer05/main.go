package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// using prefixer function
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
	// get a filename to read it's content
	args := os.Args
	if len(args) < 2 {
		log.Println("Usage: go run main.go <filename>")
		return
	}

	count, err := fileLen(args[1])
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("File length:", count)
	}

	result, err := div(10, 0)
	if err != nil {
		log.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(file, ": No such file or directory")
	}
	defer file.Close()
	data := make([]byte, 64)
	total := 0
	for {
		count, err := file.Read(data)
		if err != nil {
			if err == io.EOF {
				return total + count, nil
			}
			return 0, err
		}
		total += count
	}
}

func prefixer(str string) func(string) string {
	return func(s string) string {
		return str + " " + s
	}
}
