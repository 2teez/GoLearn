package format

import (
	"ch09/math"
	"fmt"
)

func Number[n math.Integer](num n) string {
	return fmt.Sprintf("The number is %v", num)
}
