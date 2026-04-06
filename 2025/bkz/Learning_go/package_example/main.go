package main

import (
	"fmt"
	format "package_example/do-format"
	"package_example/maths"
)

func main() {
	num := maths.Double(2)
	output := format.Number(num)
	fmt.Println(output)
}
