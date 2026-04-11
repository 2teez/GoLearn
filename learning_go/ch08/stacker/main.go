package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func main() {
	stack := &Stack[int]{}
	fmt.Println(stack)
	//
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack)
	//
	item, ok := stack.Pop()
	if ok {
		fmt.Println(item)
	}
	item, ok = stack.Pop()
	if ok {
		fmt.Println(item)
	}
	item, ok = stack.Pop()
	if ok {
		fmt.Println(item)
	}
	item, ok = stack.Pop()
	if !ok {
		fmt.Println("Stack is empty")
	}
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zero T
		return zero, false
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item, true
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.data)
}
