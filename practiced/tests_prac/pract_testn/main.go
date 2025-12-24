package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Start Here...")
}

func Divide(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Zero Division, not proper!")
	}
	return x / y, nil
}
