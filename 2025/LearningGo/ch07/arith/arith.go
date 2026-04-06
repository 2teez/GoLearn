package arith

import "fmt"

type Integer interface {
	~int | ~float32 | ~float64 | ~uint64
}

type Arithmetics[T Integer] func(a, b T) (T, error)

func Add[T Integer](first, second T) (T, error) {
	return first + second, nil
}

func Minus[T Integer](first, second T) (T, error) {
	return first - second, nil
}

func Mul[T Integer](first, second T) (T, error) {
	return first * second, nil
}

func Div[T Integer](first, second T) (T, error) {
	if second == 0 {
		return 0, fmt.Errorf("can't divide by zero.")
	}
	return first / second, nil
}
