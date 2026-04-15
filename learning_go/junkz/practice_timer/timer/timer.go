package person

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       time.Time
}

func NewPerson(firstname, lastname string, age time.Time) *Person {
	return &Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}
}

func (p *Person) CalculatedAge() int {
	now := time.Now()
	return now.Year() - p.Age.Year()
}

func (p *Person) CalculateMonthAgeAt(now time.Time) int {
	return int(now.Sub(p.Age).Hours() / 24 / 30)
}

func (p *Person) String() string {
	return fmt.Sprintf("Person[FirstName: %s, LastName: %s, Age: %d]\n",
		p.FirstName, p.LastName, p.CalculatedAge())
}
