package main

import (
	"fmt"
	"math/rand"
)

func main() {

	intStack := Stack[int]{}
	strStack := Stack[string]{}
	intStack.push(rand.Intn(50))
	intStack.push(rand.Intn(50))
	strStack.push("clojure")
	strStack.push("java")
	strStack.push("erlang")
	fmt.Println(strStack, intStack)
	fmt.Println(strStack.peek(), strStack)
	if value, ok := strStack.pop(); ok {
		fmt.Println(value)
	}
	fmt.Println(strStack.peek(), strStack)

	floatstack := Stack[float64]{}
	floatstack.cPush(3.76).cPush(1.56).cPush(9.32).cPush(1.7)
	fmt.Println(floatstack)
	///
	fmt.Println(add(1, 7, 8), add("java", "script"),
		reduce(0, []int{4, 8, 2, 1}))
	///
	value := myNumber[myInt]{value: 5}
	printIt(value)

	valuef := myNumber[myFloat]{value: 22.0 / 7}
	printIt(valuef)
}
