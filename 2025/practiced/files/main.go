package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	path, err := getUserInput("Enter your filename: ")
	if err != nil {
		log.Fatal(err)
	}
	nerr := findFile(path)
	if nerr != nil {
		log.Fatal(nerr)
	}
}

func getUserInput(msg string) (string, error) {
	fmt.Print(msg)
	buf := bufio.NewReader(os.Stdin)
	var input string
	input, err := buf.ReadString('\n')
	if err != nil {
		return input, err
	}
	return strings.TrimSpace(input), nil
}

func findFile(directory string) error {
	info, err := os.ReadDir(directory)
	if err != nil {
		return err
	}
	for _, file := range info {
		filePath := filepath.Join(directory, file.Name())
		if file.IsDir() {
			err := findFile(filePath)
			if err != nil {
				return err
			}
		} else if file.Type().IsRegular() {
			fmt.Println(filePath)
		}
	}
	return nil
}
