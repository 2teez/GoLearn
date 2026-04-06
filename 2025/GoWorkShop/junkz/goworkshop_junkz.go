package main

import (
	"fmt"
	"time"
)

func elapaseTime() {
	start := time.Now()
	time.Sleep(time.Millisecond * 2000)
	end := time.Now()
	spentTime := end.Sub(start)
	fmt.Println("The execution took exactly", spentTime, "seconds!")
}

func futureTime(length int) {
	start := time.Now()
	fmt.Println(start.Format(time.ANSIC))
	hours := start.Hour() + length
	minutes := start.Minute() + length
	seconds := start.Second() + length

	inFuture := time.Date(start.Year(), start.Month(), start.Day(),
		hours, minutes, seconds, start.Nanosecond(), time.UTC)
	fmt.Printf(
		"The current time is : %v, but in %d hours, %d minutes and %d seconds, it will be %+#v\n",
		start, length, length, length, inFuture)
}

func performATask(msg string, f func()) {
	fmt.Println(msg)
	start := time.Now()
	f()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf(
		"It takes \nHours: %v,\nMinutes: %v,\nSeconds: %v,\nMilliseconds: %v,\nNanosconds: %v\n",
		int(duration.Hours()), int(duration.Minutes()), int(duration.Seconds()),
		int(duration.Milliseconds()), int(duration.Nanoseconds()))
}

func displayPresentTime() {
	fmt.Println(time.Now().Month(), time.Now().Day(), time.Now().Weekday(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())
}
func displayTimeMinutesAndSeconds() string {
	seconds := time.Second
	minutes := time.Minute
	hour := time.Hour
	return fmt.Sprintf("%v, %v, %v", seconds, minutes, hour)
}
func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}
