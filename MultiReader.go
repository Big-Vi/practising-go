package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the first CSV file
	f1, err := os.Open("files/airtravel.csv")
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer f1.Close()

	// Open the second CSV file
	f2, err := os.Open("files/addresses.csv")
	if err != nil {
		fmt.Println("Error opening the file", err)
		return
	}
	defer f2.Close()

	// Create a MultiReader to read from both CSV files
	multiReader := io.MultiReader(f1, f2)

	// Parse the CSV files
	err = parseFiles(multiReader)
	if err != nil {
		fmt.Println(err)
	}
}

// parseFiles reads from a given io.Reader and prints CSV records
func parseFiles(mr io.Reader) error {
	// Create a CSV reader with LazyQuotes enabled
	csvReader := csv.NewReader(mr)
	csvReader.LazyQuotes = true

	// Read and process CSV records
	for {
		// Read a record from the CSV reader
		record, err := csvReader.Read()

		// Check for the end of the file
		if err == io.EOF {
			return nil
		}

		// Check for other errors
		if err != nil {
			return err
		}

		// Print the CSV record (replace this with your processing logic)
		fmt.Println(record)
	}
}
