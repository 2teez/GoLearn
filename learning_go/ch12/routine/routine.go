package routine

import (
	"fmt"
	"time"
)

type IterationParams struct {
	Start int
	Stop  int
	Step  int
}

func Summation(params IterationParams) int {
	sum := 0
	for i := params.Start; i <= params.Stop; i += params.Step {
		sum += i
	}
	return sum
}

func Routine() {
	fmt.Println("Start")
	go hello()
	time.Sleep(time.Second)
	fmt.Println("End")
}

func hello() {
	fmt.Println("Hello, World")
}
