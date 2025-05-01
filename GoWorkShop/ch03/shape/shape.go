package shape

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
	Name() string
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Name() string {
	return fmt.Sprintf("%v", reflect.TypeOf(t))
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Name() string {
	return fmt.Sprintf("%v", reflect.TypeOf(r))
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	Length float64
}

func (s Square) Name() string {
	return fmt.Sprintf("%v", reflect.TypeOf(s))
}

func (s Square) Area() float64 {
	return math.Pow(s.Length, 2.0)
}
