package shape

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type Shape interface {
	Area() float64
	Name() string
}

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Name() string {
	return getRealName(t)
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Name() string {
	return getRealName(r)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	Length float64
}

func (s Square) Name() string {
	return getRealName(s)
}

func (s Square) Area() float64 {
	return math.Pow(s.Length, 2.0)
}

func getRealName(s Shape) string {
	named := fmt.Sprintf("%v", reflect.TypeOf(s))
	if name, ok := strings.CutPrefix(named, "shape."); ok {
		named = name
	}
	return named
}
