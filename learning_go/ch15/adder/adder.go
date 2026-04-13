package adder

type Numbers interface{ ~int | ~float64 }

func addNumbers[T Numbers](a, b T) T {
	return a + b
}
