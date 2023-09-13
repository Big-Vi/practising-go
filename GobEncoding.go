package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// Student represents a student's information.
type Student struct {
	ID   string
	Name string
	Age  int
}

func main() {
	// Create a Student instance
	s := Student{
		ID:   "1",
		Name: "Vicky",
		Age:  34,
	}

	// Create a file for writing
	fileCreate, createErr := os.Create("files/student.gob")
	if createErr != nil {
		fmt.Println("Error creating a file:", createErr)
		return
	}
	defer fileCreate.Close() // Ensure the file is closed when done

	// Create a Gob encoder for the file
	enc := gob.NewEncoder(fileCreate)

	// Encode and write the Student data to the file
	if err := enc.Encode(s); err != nil {
		log.Fatal(err)
	}

	// Decode
	// Open the file for reading
	fileOpen, openErr := os.Open("files/student.gob")
	if openErr != nil {
		fmt.Println("Error opening the file:", openErr)
		return
	}
	defer fileOpen.Close() // Ensure the file is closed when done

	// Create a new Student instance for decoding
	student := Student{}

	// Create a Gob decoder for the file
	dec := gob.NewDecoder(fileOpen)

	// Decode and read the Student data from the file
	if err := dec.Decode(&student); err != nil {
		log.Fatal(err)
	}

	// Print the decoded Student data
	fmt.Println(student)
}
