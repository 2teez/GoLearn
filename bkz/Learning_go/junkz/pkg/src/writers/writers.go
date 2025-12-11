package writers

import (
	"fmt"
)

func Pp[T any](value T) {
	fmt.Println(value)
}
