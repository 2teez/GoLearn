package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"practice_fs/csvs"
	"practice_fs/files"
	"practice_fs/flags"
	"strconv"
	"strings"
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

	// using csvs package
	type Record struct {
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		Age       int    `json:"age"`
	}

	//
	csvData := csvs.Read("names.csv")
	personalData := make([]Record, 0, len(csvData)-1)
	for _, record := range csvData[1:] {
		age, _ := strconv.Atoi(strings.TrimSpace(record[2]))
		personalData = append(personalData, Record{
			Firstname: record[0],
			Lastname:  record[1],
			Age:       age,
		})
	}
	fmt.Println(personalData)
	/*
			json, err = json.MarshalIndent([]byte(`personalData`), "", " ")
		if err != nil {
			log.Fatal(err)
		}
		writer.Write(json)
	*/
}
