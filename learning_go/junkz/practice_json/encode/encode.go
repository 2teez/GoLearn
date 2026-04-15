package encode

import (
	"encoding/json"
	"log"
)

type Greeting struct {
	Message string `json:"message"`
}

func Greet(s Greeting) []byte {
	json, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return json
}
