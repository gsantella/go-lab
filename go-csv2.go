package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// Read a CSV file
	file, err := os.Open("students.csv")
	if err != nil {
		log.Fatalln("Error opening file")
	}
	defer file.Close()

	// data structure for the CSV file data
	var records [][]string

	reader := csv.NewReader(file)
	records, err = reader.ReadAll()
	if err != nil {
		log.Fatalln("Error reading CSV data")
	}

	for _, record := range records {
		fmt.Println(record)
	}

}
