package flags

import (
	"flag"
)

func Runs() (*int, *string, *bool) {
	name := flag.String("name", "", "What is your name?")
	age := flag.Int("age", -1, "what is your age?")
	married_status := flag.Bool("status", false, "Married? ")
	flag.Parse()
	return age, name, married_status
}
