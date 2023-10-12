package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("files/lorem.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read 5 bytes from the 6th byte in the file
	buf := make([]byte, 5)
	if _, err := file.Seek(6, io.SeekStart); err != nil {
		fmt.Println("Error seeking:", err)
		return
	}
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Content from 6th byte:", string(buf[:n]))

	// Read 5 bytes from the current file pointer position (after the previous read)
	if _, err := file.Seek(5, io.SeekCurrent); err != nil {
		fmt.Println("Error seeking:", err)
		return
	}
	n, err = file.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Content from current position + 5:", string(buf[:n]))

	// Read the last 5 bytes of the file
	if _, err := file.Seek(-5, io.SeekEnd); err != nil {
		fmt.Println("Error seeking:", err)
		return
	}
	n, err = file.Read(buf)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Last 5 bytes of the file:", string(buf[:n]))
}
