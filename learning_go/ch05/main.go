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
}
