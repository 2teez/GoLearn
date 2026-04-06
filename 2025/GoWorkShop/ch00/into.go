package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type String string

type Patient struct {
	FirstName     string
	LastName      string
	Age           int
	PeanutAllergy bool
}

func (p *Patient) String() string {
	return fmt.Sprintf("Patient {\n\tName: %s, %s\n\tAge: %d\n\tPeanut Allergy: %v\n}",
		p.LastName, p.FirstName, p.Age, p.PeanutAllergy)
}

func makePatient(firstname, lastname string, age int, peanutAllergy bool) *Patient {
	return &Patient{
		String(firstname).toLocalTitle(),
		String(lastname).toLocalTitle(), age, peanutAllergy,
	}
}

func makePatientDefault() *Patient {
	return &Patient{}
}

func (s String) toLocalTitle() string {
	ls := string(s)
	ls = strings.ToUpper(ls)
	return (ls[:1] + strings.ToLower(ls[1:]))
}

func random() {
	for range 10 {
		r := rand.Intn(10) + 1
		fmt.Println(r, strings.Repeat("*", r))
	}
}
