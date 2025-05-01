package main

import (
	"fmt"

	"github.com/practice/msg/msg"
	"github.com/practice/msg/shape"
)

func main() {
	fmt.Print(msg.Greeting("Hello...\n"))
	shapes := []shape.Shape{shape.Triangle{Base: 6, Height: 4.5},
		shape.Square{Length: 4.25}, shape.Rectangle{Height: 3.12, Width: 4.21}}
	printShapesDetails(shapes...)

}

func printShapesDetails(shapes ...shape.Shape) {
	for _, sh := range shapes {
		fmt.Printf("Name: %s, Area: %.2f\n", sh.Name(), sh.Area())
	}
}
