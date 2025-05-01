package main

import (
	"fmt"

	"github.com/practice/msg/msg"
	s "github.com/practice/msg/shape"
)

func main() {
	fmt.Print(msg.Greeting("Hello...\n"))
	shapes := []s.Shape{s.Triangle{Base: 6, Height: 4.5},
		s.Square{Length: 4.25}, s.Rectangle{Height: 3.12, Width: 4.21}}
	printShapesDetails(shapes...)

}

func printShapesDetails(shapes ...s.Shape) {
	for _, sh := range shapes {
		fmt.Printf("Name: %s, Area: %.2f\n", sh.Name(), sh.Area())
	}
}
