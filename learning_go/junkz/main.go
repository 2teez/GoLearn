package main

import (
	"fmt"
	person "junkz/timer"
	"time"
)

func main() {
	java := person.NewPerson("java", "james", time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC))
	clojure := person.NewPerson("clojure", "clojure", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC))
	elixir := person.NewPerson("elixir", "elixir", time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC))
	var ppl []person.Person = []person.Person{*java, *clojure, *elixir}

	for _, p := range ppl {
		fmt.Println(p.String())
	}

	// Calculate month age at a specific time
	now := time.Now()
	aperson := person.NewPerson("aperson", "aperson", time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC))
	fmt.Printf("%s is %d months old at %s\n", aperson.FirstName, aperson.CalculateMonthAgeAt(now), now)

}
