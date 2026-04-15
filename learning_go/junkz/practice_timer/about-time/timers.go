package timers

import (
	"fmt"
	"time"
)

func RunTimer() {
	start := time.Now()
	fmt.Println("The script started at", start)
	fmt.Println("Saving the World....")
	time.Sleep(time.Second * 2)
	end := time.Now()
	fmt.Println("The script ended at", end)
	fmt.Println("The script took", end.Sub(start), "to run")
}

func RunTimer2() {
	day := time.Now().Weekday()
	hour := time.Now().Hour()

	fmt.Println("Day of Week: ", day, "Hour of the day: ", hour)

	if day.String() == "Monday" {
		if hour == 1 {
			fmt.Println("Perform full-blown test..")
		} else {
			fmt.Println("Perform hit and run test..")
		}
	} else {
		fmt.Println("Perform git and run test...")
	}
}

func WhatSayTime() string {
	return time.Now().Format(time.ANSIC)
}
