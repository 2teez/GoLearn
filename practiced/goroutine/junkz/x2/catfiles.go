package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getUserInput(msg string) (string, error) {
	fmt.Print(msg)
	buf := bufio.NewReader(os.Stdin)
	input, err := buf.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(input), nil
}

func readFiles(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bytes))
}

func readFilesAndSearch(filename, text string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		found := scanner.Text()
		if strings.Contains(strings.ToLower(found), strings.ToLower(text)) {
			fmt.Printf("Search text %q found in file %s\n", text, filename)
			fmt.Printf("Line: %s\n", found)
			return
		}
	}
}
