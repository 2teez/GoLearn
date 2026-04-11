package main

import (
	"ch09/convert"
	format "ch09/do-format"
	"ch09/math"
	"fmt"
)

func main() {
	intValue := 12
	fmt.Println(intValue, format.Number(math.Double(intValue)))
	var floatValue float64 = 3.14
	fmt.Println(floatValue, format.Number(math.Double(floatValue)))

	numbers := []any{2, 7, 8, 1, 0.23}
	for _, num := range numbers {
		switch v := num.(type) {
		case float64:
			fmt.Println(format.Number(math.Double(convert.From[float64, convert.MyFloat](v))))
		case int:
			fmt.Println(format.Number(math.Double(convert.From[int, convert.MyInt](v))))
		case convert.MyInt:
			fmt.Println(format.Number(math.Double(v).To()))
		case convert.MyFloat:
			fmt.Println(format.Number(math.Double(v).To()))
		}
	}

	anum := 348.567
	fmt.Println(convert.From[float64, int](anum))

}
