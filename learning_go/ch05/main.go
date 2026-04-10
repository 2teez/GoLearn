package main

import (
	"fmt"
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
			fmt.Println(j, tables[i+1](j))
		}
	}
}
func times(base int) Times {
	return func(factor int) int {
		return base * factor
	}
}
