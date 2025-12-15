package testings

import (
	"testing"
)

func TestAddition(t *testing.T) {
	values := []int{3, 8, 0, 2, 6, 4}
	ans := []int{11, 8, 2, 8, 10}
	for ind := range len(values) - 1 {
		a, b := values[ind], values[ind+1]
		got := ans[ind]
		if Add(a, b) != got {
			t.Errorf("%v + %v = %d", a, b, a+b)
		}
	}
}

func TestSubtraction(t *testing.T) {
	values := []int{3, 8, 0, 2, 6, 4}
	ans := []int{-5, 8, -2, -4, 2}
	for ind := range len(values) - 1 {
		a, b := values[ind], values[ind+1]
		got := ans[ind]
		if Subtract(a, b) != got {
			t.Errorf("%v - %v = %d", a, b, a-b)
		}
	}
}
