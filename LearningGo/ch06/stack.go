package main

import "fmt"

type Stack[T comparable] struct {
	vals []T
}

func (s *Stack[T]) push(value T) {
	s.vals = append(s.vals, value)
}

func (s *Stack[T]) cPush(value T) *Stack[T] { // cPush is chain Push
	s.vals = append(s.vals, value)
	return s
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.vals) == 0 {
		var zero T
		return zero, false
	}
	result := s.vals[len(s.vals)-1]
	s.vals = s.vals[:len(s.vals)-1]
	return result, true
}

func (s *Stack[T]) peek() T {
	if len(s.vals) == 0 {
		var zero T
		return zero
	}
	result := s.vals[len(s.vals)-1]
	return result
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("Stack{%p}", s.vals)
}
