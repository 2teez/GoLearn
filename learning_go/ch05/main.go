package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

type Times func(int) int

func main() {
	people := []person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Println(people)
	sort.Slice(people, func(a, b int) bool {
		return people[a].Age < people[b].Age
	})
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].FirstName < people[j].FirstName
	})
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// function returning functions
	// make times table bases
	tables := make([]Times, 0, 3)
	for i := 2; i < 5; i++ {
		tables = append(tables, times(i))
	}

	for i := range len(tables) {
		fmt.Println("Times table for ", i+2, ":")
		for j := 2; j <= 12; j++ {
			fmt.Println(j, tables[i](j))
		}
	}

	// read file
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage args[0] <filename>")
	}
	readFile(args[1])
}
func times(base int) Times {
	return func(factor int) int {
		return base * factor
	}
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	data := make([]byte, 2048)
	for {
		n, err := file.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		fmt.Println(string(data[:n]))
	}

}
