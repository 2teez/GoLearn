package routine

import (
	"fmt"
)

func Routine() {
	fmt.Println("Start")
	go hello()
	fmt.Println("End")
}

func hello() {
	fmt.Println("Hello, World")
}
