package main

import (
	"encoding/json"
	"fmt"
	"strings"

	d "practicejson/decoder"
	e "practicejson/encoder"
)

func main() {
	names := []string{"java", "python", "go"}
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(names)
	fmt.Println(writer.String())

	e.RunEncodeDefaultTypes()
	e.RunEncodeCustomTypes()
	e.RunProductCustomTypes()
	d.DecodeConfigFile("config.json")
	/*
		 	jsonData, err := json.Marshal(names)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(jsonData))
	*/
}
