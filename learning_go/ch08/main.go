package main

import (
	"fmt"
	"strings"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 |
		~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type Printable interface {
	~int | ~float64
	fmt.Stringer
}

type MyInt int
type MyFloat float64

type linkedList[T any] struct {
	val  T
	next *linkedList[T]
}

func main() {
	fmt.Println(double(34.5))
	fmt.Println(double(0.054))
	fmt.Println(double(3))
	//
	var intValue MyInt = 3
	printIt(intValue)
	printIt(MyFloat(3.14))
	//
	// linked list
	list := &linkedList[int]{val: 1}
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	fmt.Printf("%v\n", list)
	list.insert(6, 2)
	fmt.Printf("%v\n", list)
	list.remove(2)
	fmt.Printf("%v\n", list)
}

func (l *linkedList[T]) Add(val T) {
	newNode := l
	for newNode.next != nil {
		newNode = newNode.next
	}
	newNode.next = &linkedList[T]{val: val}
}

func (l *linkedList[T]) insert(val T, index int) {
	newNode := l
	for i := 0; i < index-1; i++ {
		newNode = newNode.next
	}
	newNode.next = &linkedList[T]{val: val, next: newNode.next}
}

func (l *linkedList[T]) remove(index int) {
	newNode := l
	for i := 0; i < index-1; i++ {
		newNode = newNode.next
	}
	newNode.next = newNode.next.next
}

func (l *linkedList[T]) String() string {
	newNode := l
	result := strings.Builder{}
	for newNode != nil {
		fmt.Fprintf(&result, "%v -> ", newNode.val)
		newNode = newNode.next
	}
	result.WriteString("nil")
	return result.String()
}

func double[I Integer](i I) I {
	return i + i
}

func printIt[P Printable](p P) {
	fmt.Println(p)
}

func (i MyInt) String() string {
	return fmt.Sprintf("%v", int(i))
}

func (f MyFloat) String() string {
	return fmt.Sprintf("%v", float64(f))
}
