package main

import (
	r "ch12/routine"
	"fmt"
	"time"
)

func main() {

	r.Routine()
	oneToTen := r.IterationParams{Start: 1, Stop: 10, Step: 1}
	oneToHundred := r.IterationParams{Start: 1, Stop: 100, Step: 1}

	{ // using goroutines with time.Sleep
		var sumResult1, sumResult2 int
		go func() {
			sumResult1 = r.Summation(oneToTen)
		}()
		go func() {
			sumResult2 = r.Summation(oneToHundred)
		}()
		time.Sleep(time.Second * 1)
		fmt.Println(sumResult1, sumResult2)
	}
}
