package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	add   = func(i, j int) (int, error) { return i + j, nil }
	minus = func(i, j int) (int, error) { return i - j, nil }
	times = func(i, j int) (int, error) { return i * j, nil }
	div   = func(i, j int) (int, error) {
		if j == 0 {
			return 0, errors.New("Can't divide by zero.")
		}
		return i / j, nil
	}
)

func main() {
	if _, err := div(4, 0); err != nil {
		log.Println(err)
	}
	var filename string
	fmt.Print("Enter your filename: ")
	fmt.Scanln(&filename)
	fmt.Println(fileLen(filename))
	//
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}

func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		if err != io.EOF {
			return 0, err
		}
	}
	defer f.Close()
	bytes := make([]byte, 2048)
	n, err := f.Read(bytes)
	return n, err
}

func prefixer(greet string) func(string) string {
	return func(who string) string {
		return fmt.Sprintf("%v %v", greet, who)
	}
}
