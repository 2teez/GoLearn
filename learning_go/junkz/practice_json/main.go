package main

import (
	"fmt"
	"practice_json/decode"
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
}
