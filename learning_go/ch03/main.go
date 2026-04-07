package main

import (
	"fmt"
)

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	firstSlice := greetings[:2]
	secondSlice := greetings[2:4]
	thirdSlice := greetings[3:]
	fmt.Println(firstSlice, secondSlice, thirdSlice)

	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	firstEmployee := Employee{}
	secondEmployee := Employee{firstName: "John", lastName: "Doe", id: 1}
	var thirdEmployee Employee
	thirdEmployee.firstName = "Jane"
	thirdEmployee.lastName = "Doe"
	thirdEmployee.id = 2
	fmt.Println(firstEmployee, secondEmployee, thirdEmployee)

}
