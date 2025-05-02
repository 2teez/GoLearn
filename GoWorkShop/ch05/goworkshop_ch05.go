package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(log.Ldate | log.Llongfile)
}

func createFile(filename string) (*os.File, error) {
	f, err := os.Create(filename)

	if err != nil {
		return nil, err
	}
	return f, nil
}

func writeToFile(file *os.File, s string) bool {
	_, err := file.WriteString(s)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	return true
}

func openAndReadFile(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	f, err := os.Open(filename)
	defer f.Close()
	if content, err := os.ReadFile(f.Name()); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(string(content))
	}
}

func readCSVFile(filename string) {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	f := csv.NewReader(strings.NewReader(filename))

	for {
		s, err := f.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
	}
}
