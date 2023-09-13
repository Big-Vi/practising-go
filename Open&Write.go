package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a file
	fmt.Println("Creating a file...")
	file, createErr := os.Create("files/file.txt")
	if createErr != nil {
		fmt.Println("Error creating a file:", createErr)
		return
	}
	defer file.Close() // Ensure the file is closed when done

	// Write data to the file
	fmt.Println("Writing to the file...")
	dataToRead := []byte("myfile content")
	bytesWritten, writeErr := file.Write(dataToRead)
	if writeErr != nil {
		fmt.Println("Error writing to the file:", writeErr)
		return
	}

	// Print the number of bytes written
	fmt.Printf("Number of bytes written: %d\n", bytesWritten)
}
