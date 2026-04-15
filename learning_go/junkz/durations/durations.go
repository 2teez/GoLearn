package durations

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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

func ElapsedTime(start time.Time, end time.Time) string {
	duration := end.Sub(start)
	return fmt.Sprintf("This task took: %v hours, %v minutes, %v seconds",
		strconv.Itoa(int(duration.Hours())),
		strconv.Itoa(int(duration.Minutes())),
		strconv.Itoa(int(duration.Seconds())))
}

func GetTimeAndLocation(location string) (string, error) {
	presentTime := time.Now()
	if strings.Contains(location, "/") {
		located, err := time.LoadLocation(location)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%v ~> %s is %v", presentTime.Format(time.ANSIC),
			location, presentTime.In(located).Format(time.ANSIC)), nil
	}
	return "", errors.New("Location not recognized.")
}
