package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Student represents a student's information.
type Student struct {
	Username string
	Age      int
}

func main() {
	// Create the first file
	f1, err := os.Create("files/multifile1.txt")
	if err != nil {
		fmt.Println("Error creating a file", err)
		return
	}
	defer f1.Close() // Ensure the file is closed when done

	// Create the second file
	f2, err := os.Create("files/multifile2.txt")
	if err != nil {
		fmt.Println("Error creating a file", err)
		return
	}
	defer f2.Close() // Ensure the file is closed when done

	// Sample student data
	students := []Student{
		{Username: "Vicky", Age: 33},
		{Username: "Vignesh", Age: 34},
	}

	// Use io.MultiWriter to write to both files simultaneously
	multiWriter := io.MultiWriter(f1, f2)

	// Create a JSON encoder for writing JSON data
	encoder := json.NewEncoder(multiWriter)

	// Iterate over students and encode them into the files
	for _, student := range students {
		if err := encoder.Encode(student); err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
	}

	fmt.Println("Data written to files successfully.")
}
