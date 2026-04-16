package csvs

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func Read(filename string) [][]string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(strings.NewReader(string(content)))
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Println(err)
	}
	return records
}
