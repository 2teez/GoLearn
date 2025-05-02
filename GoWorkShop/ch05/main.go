package main

import (
	"flag"
	"log"
)

func main() {
	filename := flag.String("filename", "filename", "set the file name to use")
	flag.Parse()
	file, err := createFile(*filename)
	if err != nil {
		log.Fatalln(err)
	}
	if ok := writeToFile(file, "Saying write to the file now"); ok {
		log.Println("File written and then closed.")
	}

	openAndReadFile("main.go")
	readCSVFile("Hospital Visits.csv")
	openAndReadFile("Hospital Visits.csv")
}
