package main

import (
	"fmt"
	"practice_json/decode"
	"practice_json/encode"
	execr "practice_json/exer"
)

func main() {
	data := []byte(`
	{
        	"message": "Greetings fellow gopher..."
	}`)

	var greet decode.Greeting
	fmt.Println(decode.Greet(data, &greet))
	fmt.Println(greet.Greetings(data))

	person := []byte(`{
		"first_name": "java",
		"last_name": "gosling",
		"address": {
			"street": "Sulphur Springs Rd",
			"city": "Park City",
			"state": "VA",
			"zip": "22151"
		}
	}`)

	personData := decode.Person{}

	decode.PersonDetails(person, &personData)
	fmt.Println(personData)

	studentData, err := decode.ReadJsonFile("stud.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	student := decode.Student{}
	err = decode.ParseJson(studentData, &student)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(student)

	// Encoding
	greets := encode.Greeting{Message: "Greeting from Perl!"}
	encoded := encode.Greet(greets)
	fmt.Println(string(encoded))

	// exercise
	custData, err := decode.ReadJsonFile("cust.json")
	fmt.Println(string(custData))
	if err != nil {
		fmt.Println(err)
		return
	}

	cust := execr.Customer{}
	err = execr.ParseJson(custData, &cust)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cust)
}
