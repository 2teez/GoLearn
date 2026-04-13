package adder

import "testing"

func TestAddNumbers(t *testing.T) {
	t.Parallel()
	result := addNumbers(1, 2)
	if result != 3 {
		t.Errorf("Expected 3, got %d", result)
	}
}

func TestAddNumbersFloat(t *testing.T) {
	t.Parallel()
	result := addNumbers(1.5, 2.5)
	if result != 4.0 {
		t.Errorf("Expected 4.0, got %f", result)
	}
}
