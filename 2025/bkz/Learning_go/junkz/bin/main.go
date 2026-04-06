package main

import (
	"bufio"
	"fmt"
	calendar "junkz/pkg/src/calendar"
	"junkz/pkg/src/events"
	"junkz/pkg/src/readers"
	wrt "junkz/pkg/src/writers"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a line: ")
	n, err := readers.ReadLine()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(n)
	// using ReadString from bufio
	wrt.Pp("Using ReadString from bufio package...")
	anotherReader := bufio.NewReader(os.Stdin)
	if line, err := anotherReader.ReadString('\n'); err == nil {
		wrt.Pp(line)
	} else {
		log.Println(err)
	}

	// check the getfloat function
	age, err := readers.GetFloat("Enter your age: ")
	if err != nil {
		defer func() {
			recover()
			wrt.Pp("Get me someone's age?!")
		}()
		log.Panic(err)
	}
	wrt.Pp(age)
	// use read a file
	nerr := readers.ReadAFile("bin/main.go", func(s string) {
		fmt.Println(strings.ToUpper(s))
	})
	if nerr != nil {
		log.Fatal(nerr)
	}
	// using calender local package
	date, err := calendar.NewDate(1960, 10, 22)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", date)

	event, err := events.NewEvent("Mum's Birthday", &calendar.Date{})
	if err != nil {
		log.Fatal(err)
	}
	event.SetYear(2026)
	event.SetMonth(calendar.Aug)
	event.SetDay(13)

	fmt.Printf("%s\n", event)
}
