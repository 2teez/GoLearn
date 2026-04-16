package main

import (
	"flag"
	"fmt"
	"log"
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

}
