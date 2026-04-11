package convert

type MyInt int
type MyFloat float64

type NumberType interface {
	MyInt | MyFloat | int | float64
}

func (i MyInt) To() int {
	return int(i)
}

func (f MyFloat) To() float64 {
	return float64(f)
}

func From[T, R NumberType](input T) R {
	return R(input)
}
