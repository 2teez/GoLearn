package decode

import (
	"encoding/json"
	"log"
)

type Greeting struct {
	Message string `json:"message"`
}

type Person struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Address   Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func (g *Greeting) Greetings(data []byte) string {
	err := json.Unmarshal(data, g)
	if err != nil {
		log.Println(err)
	}
	return g.Message
}

func Greet(data []byte, greeting *Greeting) string {
	err := json.Unmarshal(data, greeting)
	if err != nil {
		log.Println(err)
	}
	return greeting.Message
}

func PersonDetails(data []byte, person *Person) {
	err := json.Unmarshal(data, person)
	if err != nil {
		log.Println(err)
	}
}
