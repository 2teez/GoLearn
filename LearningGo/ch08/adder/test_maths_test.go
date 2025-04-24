package adder

import (
	"testing"
)

func Test_addNumber(t *testing.T) {
	result, got := add_number(4, 5), 9

	if result != got {
		t.Errorf("incorrect result. Expected: %d, got: %d", 5, result)
	}
}
