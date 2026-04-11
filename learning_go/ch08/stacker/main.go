package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func main() {
	stack := &Stack[int]{}
	fmt.Println(stack)
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.data)
}
