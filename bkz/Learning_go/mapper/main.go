package main

import (
	"fmt"
	"log"
	"maps"
)

type person struct {
	name string
	age  int
}

func main() {
	langs := map[string]int{}
	fmt.Println(langs)
	langs["java"] = 32
	langs["c"] = 45
	langs["clojure"] = 12
	langs["elixir"] = 10
	fmt.Println(langs)
	if v, ok := langs["c++"]; ok {
		log.Println(v)
	} else {
		log.Printf("%v doesn't dey!", "c++")
	}
	//  using map as set
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	intSet := map[int][]interface{}{}
	for _, v := range vals {
		intSet[v] = []interface{}{}
	}
	for k := range maps.Keys(intSet) {
		fmt.Print(k, ", ")
	}
	fmt.Println()
	// make a person using person literal
	literalPerson := new(person)
	fmt.Println(literalPerson)
	literalPerson.name = "java"
	literalPerson.age = 35
	fmt.Println(literalPerson)
	//
	clojure := person{
		age:  12,
		name: "Clojure",
	}
	clojure.Print()
	literalPerson.Print()
}

func (p *person) New(name string, age int) *person {
	return &person{name, age}
}

func (p *person) String() string {
	return fmt.Sprintf("Person [Name: %v, Age: %d]\n", p.name, p.age)
}

func (p person) Print() {
	fmt.Println(&p)
}
