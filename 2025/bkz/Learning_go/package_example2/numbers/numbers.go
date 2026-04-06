package numbers

import "fmt"

type Number interface {
	~int | ~uint | ~float32 | ~float64
}

func Add[T Number](first, second T) T {
	return first + second
}

func Subtact[T Number](first, second T) T {
	return first - second
}

func Divide[T Number](first, second T) (T, error) {
	if second == 0 {
		var zero T
		return zero, fmt.Errorf("can't divide with zero!")
	}
	return first / second, nil
}
