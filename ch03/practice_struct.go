package main

import "fmt"

type person struct {
	name string
	age  int
}

func practiceStruct() {
	var julia person
	bob := person{
		"bob",
		25,
	}
	kunle := person{
		age:  34,
		name: "kunle",
	}
	julia.age = 12
	fmt.Println(julia, bob, kunle)
	fmt.Println(struct {
		name string
		age  int
	}{name: "java", age: 34})

	message := "Hi and "[3:4]
	fmt.Println([]rune(message), []byte(message), message)
	/// using a for loop & switch in golang
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word")
		case 5:
			fmt.Println(word, "is the exact word with length: ", size)
		case 6, 7, 8, 9:
			fmt.Println(word, "is a long word")
		default:
			fmt.Println(word, "too long a word to call.")
		}
	}
}
