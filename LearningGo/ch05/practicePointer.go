package main

import (
	"fmt"
)

type person struct {
	name       string
	middleName *string
	age        int
}

func (p person) String() string {
	return fmt.Sprintf("%s, %s %d", p.name, *p.middleName, p.age)
}

func main() {
	num := 34
	doubleUp(&num)
	fmt.Println(num)
	fmt.Println(doubleUpFunc(num, doubleUp))
	///
	mname := "java"
	p := person{"gosling", &mname, 34}
	fmt.Println(p)
	///
	ap := makePerson(p)
	*ap.middleName = "james"
	fmt.Println(ap, *ap.middleName, *p.middleName)
	ap2 := makePersonPointer("java", "gosling", 67)
	fmt.Println(ap2)
	//
	teams := []Team{Team{}, Team{}, Team{}}
	league := League{teams, map[string]wins{"": wins(2)}}
	fmt.Println(league)
}

func makePersonPointer(name, mname string, age int) *person {
	return &person{name, &mname, age}
}

func makePerson(p struct {
	name       string
	middleName *string
	age        int
}) person {
	return person{
		name:       p.name,
		middleName: *&p.middleName,
		age:        p.age,
	}
}

func doubleUpFunc(value int, f func(*int)) int {
	f(&value)
	return value
}

func doubleUp(value *int) {
	*value *= 2
}
