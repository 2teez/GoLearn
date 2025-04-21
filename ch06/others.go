package main

type number interface {
	~int | ~int64 | ~int32 | ~float32 | ~float64 | ~string
}

func add[T number](s T, o ...T) T {
	var result T = s
	for _, val := range o {
		result = result + val
	}
	return result
}

func reduce[T number](s T, o []T) T {
	var result T = s
	for _, val := range o {
		result += val
	}
	return result
}
