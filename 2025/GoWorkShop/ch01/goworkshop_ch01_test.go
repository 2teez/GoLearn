package main

import "testing"

func TestFizzBuzzCase15(t *testing.T) {
	expected, got := fizzBuzz(15), "FizzBuzz"
	if expected != got {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

func TestFizzBuzzCase5(t *testing.T) {
	expected, got := fizzBuzz(5), "Buzz"
	if expected != got {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

func TestFizzBuzzCase3(t *testing.T) {
	expected, got := fizzBuzz(3), "Fuzz"
	if expected != got {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

func TestFizzBuzzCase47(t *testing.T) {
	expected, got := fizzBuzz(47), "47"
	if expected != got {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}
