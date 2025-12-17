package main

import (
	"fmt"
	"time"
)

func doWork(num int) {
	fmt.Printf("Work %d started at %s\n", num, time.Now().Format("15:04:05"))
	time.Sleep(1 * time.Second)
	fmt.Printf("Work %d ended at %s\n", num, time.Now().Format("15:04:05"))
}
