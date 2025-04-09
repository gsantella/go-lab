package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {

	// Create a CSV file
	file, err := os.Create("students.csv")
	if err != nil {
		log.Fatalln("Error creating file")
	}
	defer file.Close()

	//file.WriteString("LALALALALALALALA")

	// Write some data to the CSV file
	data := [][]string{
		{"Name", "Age", "City"},
		{"Alice", "30", "New York"},
		{"Bob", "25", "Los Angeles"},
		{"Charlie", "35", "Chicago"},
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, record := range data {
		writer.Write(record)
	}

}
