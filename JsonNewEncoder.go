package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// Student represents a student with a name and age.
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Open or create a file for writing JSON data
	filePath := "files/json.txt"
	openFile, createErr := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0600)
	if createErr != nil {
		fmt.Printf("Error opening file: %v", createErr)
		return
	}
	defer openFile.Close() // Ensure the file is closed when done

	// Create a buffer to hold the JSON data
	buf := new(bytes.Buffer)

	// Encode a Student struct as JSON and write it to the buffer
	encoder := json.NewEncoder(buf)
	s := Student{Name: "Vignesh", Age: 33}
	if enocdeErr := encoder.Encode(s); enocdeErr != nil {
		fmt.Printf("Error encoding JSON: %v", enocdeErr)
		return
	}

	// Write the JSON data from the buffer to the file
	openFile.Write(buf.Bytes())
}
