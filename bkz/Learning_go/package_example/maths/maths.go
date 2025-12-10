package maths

type Integer interface {
	~int | ~uint | ~float32 | ~float64
}

func Double[T Integer](val T) T {
	return val * 2
}
