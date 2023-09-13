package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Create a string reader with content
	contentReader := strings.NewReader("file content")

	// Create a file for writing
	outputFile, createErr := os.Create("files/file.txt")
	if createErr != nil {
		fmt.Printf("Error creating a file: %v", createErr)
		return
	}
	defer outputFile.Close() // Ensure the file is closed when done

	// Use io.TeeReader to read from the content reader and write to the output file simultaneously
	teeReader := io.TeeReader(contentReader, outputFile)

	// Read all data from the TeeReader
	dataBytes, readErr := io.ReadAll(teeReader)
	if readErr != nil {
		fmt.Printf("Error reading: %v", readErr)
		return
	}

	// Print the data
	fmt.Println(string(dataBytes))
}
