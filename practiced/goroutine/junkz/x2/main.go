package main

import (
	"log"
	"os"
	"time"
)

func main() {
	const times int = 5
	for i := range times {
		go doWork(i)
	}
	time.Sleep(2 * time.Second)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("call the cli with more than one file...")
	}
	for _, file := range args[1:] {
		go readFiles(file)
	}
	time.Sleep(2 * time.Second)
	//
	// read files and search...
	input, err := getUserInput("Enter the string to search: ")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range args[1:] {
		go readFilesAndSearch(file, input)
	}
	time.Sleep(2 * time.Second)
}
