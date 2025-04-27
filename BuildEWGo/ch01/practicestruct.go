package main

import "fmt"

type Greeter interface {
	sayHello() string
}

type Person struct {
	name string
	age  int
}

func (p Person) sayHello() string {
	return fmt.Sprintf("Hi! my name is %s.", p.name)
}

func makePerson(name string, age int) *Person {
	return &Person{name, age}
}
