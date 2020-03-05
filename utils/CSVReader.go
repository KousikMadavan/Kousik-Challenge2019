package utils

import (
	"encoding/csv"
	"log"
	"os"
)

// CommonCSVReader ...
func CommonCSVReader(filename string) ([][]string, error) {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Couldn't read content from CSV file", err)
	}
	return records, nil
}
