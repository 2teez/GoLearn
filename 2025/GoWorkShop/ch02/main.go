package main

import (
	"fmt"
	"reflect"
)

func main() {
	b := makeDirectDeposite(
		struct {
			firstName, lastName          string
			routingNumber, accountNumber int
		}{"smith", "", 87, 1809})
	b.validateRoutingNumber()
	b.validateLastName()
	fmt.Println(reflect.TypeOf(InvalidLastName))
}
