package main

import (
	"bufio"
	"fmt"
	"junkz/pkg/src/readers"
	wrt "junkz/pkg/src/writers"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a line: ")
	n, err := readers.ReadLine()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(n)
	// using ReadString from bufio
	wrt.Pp("Using ReadString from bufio package...")
	anotherReader := bufio.NewReader(os.Stdin)
	if line, err := anotherReader.ReadString('\n'); err == nil {
		wrt.Pp(line)
	} else {
		log.Println(err)
	}

	// check the getfloat function
	age, err := readers.GetFloat("Enter your age: ")
	if err != nil {
		defer func() {
			recover()
			wrt.Pp("Get me someone's age?!")
		}()
		log.Panic(err)
	}
	wrt.Pp(age)
	// use read a file
	nerr := readers.ReadAFile("bin/main.go", func(s string) {
		fmt.Println(strings.ToUpper(s))
	})
	if nerr != nil {
		log.Fatal(nerr)
	}
}
