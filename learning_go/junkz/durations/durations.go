package durations

import (
	"fmt"
	"time"
)

func DurationsRun() {
	start := time.Now()
	fmt.Println("The script started at: ", start)
	sum := 0
	for i := 0; i <= 10_000_000_000; i++ {
		sum += i
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("The script is completed at: ", end)
	fmt.Println("This task took: ", duration.Hours(), "hours")
	fmt.Println("This task took: ", duration.Minutes(), "minutes")
	fmt.Println("This task took: ", duration.Seconds(), "seconds")
	fmt.Println("This task took: ", duration.Nanoseconds(), "nanoseconds")
}

func AddSeconds(duration time.Time, seconds int) time.Time {
	return duration.Add(time.Second * time.Duration(seconds))
}

func AddMinutes(duration time.Time, minutes int) time.Time {
	return duration.Add(time.Minute * time.Duration(minutes))
}

func AddHours(duration time.Time, hours int) time.Time {
	return duration.Add(time.Hour * time.Duration(hours))
}
