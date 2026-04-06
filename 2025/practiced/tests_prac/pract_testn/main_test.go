package main

import (
	"testing"
)

func TestDivisionFunction(t *testing.T) {
	t.Parallel()
	x, y := 3, 7
	got := checker(t, x, y)
	want := 0
	if got != want {
		t.Errorf("Maths doesn't work Again!...")
	}
}

func checker(t *testing.T, firstVal, secondVal int) int {
	t.Helper()
	got, err := Divide(firstVal, secondVal)
	if err != nil {
		t.Fatal(err)
	}
	return got
}
