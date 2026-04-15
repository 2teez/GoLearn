package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	//fmt.Println(string(custData))
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

	// initial customer's purchase order
	startedOrder, err := json.MarshalIndent(cust, "", " ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(startedOrder))

	orders := []execr.Item{
		execr.Item{Name: "Final Fantasy The Zodiac Age", Description: "Nintendo Switch Game", Quantity: 1, Price: 50.0},
		execr.Item{Name: "Crystal Drinking Glass", Quantity: 11, Price: 25.0},
	}

	// add orders to the customer's purchase order
	cust.PurchaseOrder.OrderDetail = append(cust.PurchaseOrder.OrderDetail, orders...)

	// calculate total price
	for _, item := range cust.PurchaseOrder.OrderDetail {
		cust.PurchaseOrder.TotalPrice += item.Price * float64(item.Quantity)
	}

	// mark order as paid
	cust.PurchaseOrder.IsPaid = true

	// get the original JSON data
	original, err := json.MarshalIndent(cust, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(original))
}
