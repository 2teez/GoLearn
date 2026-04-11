package math

type Integer interface {
	~int | ~float64
}

func Double[T Integer](a T) T {
	return a * 2
}
