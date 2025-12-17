package main

import "time"

func main() {
	const times int = 5
	for i := range times {
		go doWork(i)
	}
	time.Sleep(2 * time.Second)
}
