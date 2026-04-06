package main

import "fmt"

func main() {
	var v greeting
	data := []byte(`
		{
		"message": "Hello, Gophers!"
		}
		`)
	unmarshalJson(data, &v)
	fmt.Printf("%s, %+v\n", v.Message, v)
	var student student
	stdData := []byte(`{
		"id": 123,
		"lname": "Smith",
		"minitial": null,
		"fname": "John",
		"enrolled": true,
		"classes": [{
		"coursename": "Intro to Golang",
		"coursenum": 101,
		"coursehours": 4
		},
		{
		"coursename": "English Lit",
		"coursenum": 101,
		"coursehours": 3
		},
		{
		"coursename": "World History",
		"coursenum": 101,
		"coursehours": 3
		}
		]
		}
`)
	unmarshalJson(stdData, &student)
	fmt.Printf("%+v", student)

	//
	marshalJson(student)
	marshalJson(v)
}
