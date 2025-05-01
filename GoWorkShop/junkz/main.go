package main

import (
	"fmt"
	"time"
)

func main() {
	// practicing time in golang
	fmt.Println(whatstheclock())
	fmt.Println(displayTimeMinutesAndSeconds())
	displayPresentTime()
	performATask("Take a sum of a 100000 ints using traditional for-loop", func() {
		sum := 0
		for i := 0; i < 100_001; i++ {
			sum += i
		}
		time.Sleep(2 * time.Second)
	})

	performATask("\nTake a sum of a 100000 ints using for-loop/range", func() {
		sum := 0
		for i := range 100_001 {
			sum += i
		}
		time.Sleep(2 * time.Second)
	})

	_dob, _ := time.Parse(time.RFC3339, "1978-02-05T22:0:0+00:00")
	fmt.Println(_dob)
	dob := time.Date(1960, 2, 5, 0, 0, 0, 12345, time.UTC)
	fmt.Println(time.Now().Sub(dob), time.Now().Year()-dob.Year()) // age in hours,mins,secons,millis
	//
	currentTime := time.Now()
	locations := map[string]string{
		"tokyo":    "Asia/Tokyo",
		"berlin":   "Europe/Berlin",
		"lagos":    "Africa/Lagos",
		"new_york": "America/New_York",
		"brisbane": "Australia/Brisbane",
		"dundee":   "Europe/London",
	}
	for key, value := range locations {
		if loc, err := time.LoadLocation(value); err == nil {
			fmt.Printf("%10s %v\n", key, currentTime.In(loc).Format(time.ANSIC))
		} else {
			defer func() {
				fmt.Println(err)
			}()
		}
	}
	//
	elapaseTime()
	futureTime(struct{ len int }{3}.len)
}
