package main

import "fmt"

func main() {
	random()
	// make patient
	bob := makePatient("bob", "smith", 34, false)
	fmt.Println(bob)
	bob2 := makePatientDefault()
	fmt.Println(bob2)
}
