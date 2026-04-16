package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"practice_fs/files"
	"practice_fs/flags"
)

func main() {
	age, name, married_status := flags.Runs()
	if *name == "" {
		log.Println("Name is required")
		flag.PrintDefaults()
		return
	}
	fmt.Printf("Name: %s, Age: %d, Married: %t\n", *name, *age, *married_status)

	// using the file package
	writer, close, err := files.Create("temp_file.json")
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	data := struct {
		Name          string `json:"name"`
		Age           int    `json:"age"`
		MarriedStatus bool   `json:"married_status"`
	}{
		Name:          *name,
		Age:           *age,
		MarriedStatus: *married_status,
	}
	json, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	writer.Write(json)
}
