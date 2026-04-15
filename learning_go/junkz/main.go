package main

import (
	"fmt"
	timers "junkz/about-time"
	"junkz/durations"
	person "junkz/timer"
	"time"
)

func main() {
	// struct with time
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

	// Practice with Timer
	timers.RunTimer()
	timers.RunTimer2()
	fmt.Println(timers.WhatSayTime())
	nower := time.Now()
	time_parsed, _ := time.Parse(time.RFC3339, "2026-01-01T00:00:00Z")
	fmt.Println(nower, "\nTime Parsed: ", time_parsed,
		"\nDiff: ", nower.Sub(time_parsed), "\nAfter: ", nower.After(time_parsed),
		"\nBefore: ", nower.Before(time_parsed), "\nEqual: ", nower.Equal(time_parsed),
	)

	// working with durations
	durations.DurationsRun()
	timeNow := time.Now()
	fmt.Println(timeNow, "\tAfter 10 seconds: ", durations.AddSeconds(timeNow, 10))
	fmt.Println(timeNow, "\tAfter 5 minute: ", durations.AddMinutes(timeNow, 5))
	fmt.Println(timeNow, "\tAfter 1 hour: ", durations.AddHours(timeNow, 1))
}
