package decoder

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Username           string    `json:"Username"`
	AdditionalProducts []Product `json:"AdditionalProducts"`
}

type Product struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

func DecodeConfigFile(filename string) {
	strs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	var config Config
	if err := json.Unmarshal(strs, &config); err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config)
}
