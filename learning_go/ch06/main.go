package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p := MakePerson("John", "Doe", 30)
	fmt.Println(p)

	p2 := MakePersonPointer("Jane", "Doe", 25)
	fmt.Println(p2)
	///
	langs := []string{"Go", "Python", "Java"}
	fmt.Println("Before Any Functions: ", langs)
	UpdateSlice(langs, "C++")
	fmt.Println("After UpdateSlice: ", langs)
	GrowSlice(langs, "Rust")
	fmt.Println("After GrowSlice: ", langs)
}

func UpdateSlice(strs []string, str string) {
	index := len(strs) - 1
	strs[index] = str
	fmt.Println("Inside UpdateSlice:", strs)
}

func GrowSlice(strs []string, str string) {
	strs = append(strs, str)
	fmt.Println("Inside GrowSlice:", strs)
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
