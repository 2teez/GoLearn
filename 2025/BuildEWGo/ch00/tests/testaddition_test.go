package main

import (
	"strconv"
	"testing"

	"github.com/practice/addition/addition"
)

func converter(values []int) []string {
	buff := make([]string, len(values))
	for ind, value := range values {
		buff[ind] = strconv.Itoa(value)
	}
	return buff
}

func TestAdditionfunc(t *testing.T) {
	values := map[int][]int{
		37: []int{3, 7, 9, 10, 8},
		14: []int{9, 2, 3},
		17: []int{5, 3, 4, 6, -1},
	}
	for got, values := range values {
		expected := addition.Sum(0, converter(values))
		if got != expected {
			t.Errorf("Err: expected: %v, got: %v", expected, got)
		}
	}
}
