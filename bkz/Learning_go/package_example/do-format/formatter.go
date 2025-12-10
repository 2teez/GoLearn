package format

import (
	"fmt"
)

func Number[T any](num T) string {
	return fmt.Sprintf("The number is %v\n", num)
}
