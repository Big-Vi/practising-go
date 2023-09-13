package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the file for reading
	openFile, openErr := os.Open("files/lorem.txt")
	if openErr != nil {
		fmt.Printf("Error opening file: %v", openErr)
		return
	}
	defer openFile.Close() // Close the file when we're done

	buffer := make([]byte, 1024) // Create a buffer to read into

	for {
		// Read from the file into the buffer
		bytesRead, readErr := openFile.Read(buffer)
		if readErr == io.EOF {
			break // End of file reached
		} else if readErr != nil {
			fmt.Printf("Error reading into buffer: %v", readErr)
			break
		}

		// Print the data read from the file
		fmt.Println(string(buffer[:bytesRead]))
	}
}
