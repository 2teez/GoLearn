package timers

import (
	"fmt"
	"time"
)

func RunTimer() {
	start := time.Now()
	fmt.Println("The script started at", start)
	fmt.Println("Saving the World....")
	time.Sleep(time.Second * 5)
	end := time.Now()
	fmt.Println("The script ended at", end)
	fmt.Println("The script took", end.Sub(start), "to run")
}
