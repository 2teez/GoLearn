package main

import (
	"testing"
)

func TestFibonacciNumber6(t *testing.T) {
	num := 6
	ques := Fibo(num)
	got := 13
	if ques != got {
		t.Errorf("Fibo(%d) id not %d", num, got)
	}
}
