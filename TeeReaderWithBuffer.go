package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// Create a string reader with content
	content := strings.NewReader("file content")

	// Create a buffer to tee the content
	buf := new(bytes.Buffer)
	teeReader := io.TeeReader(content, buf)

	// Read all data from the TeeReader
	dataBytes, readErr := io.ReadAll(teeReader)
	if readErr != nil {
		fmt.Printf("Error reading: %v", readErr)
		return
	}

	// Print the data
	fmt.Println(string(dataBytes))
}
