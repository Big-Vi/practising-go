package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type ErrorCountWriter struct {
	Writer   io.Writer
	ErrorFile io.Writer
}

func (w *ErrorCountWriter) Write(p []byte) (n int, err error) {
	// Write the data to the original writer
	n, err = w.Writer.Write(p)

	return n, err
}

func main() {
	// Open the log file for reading
	logFilePath := "files/logfile.txt"
	logFile, err := os.Open(logFilePath)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer logFile.Close()

	// Create a custom writer that logs errors to a separate file
	errorFilePath := "files/404errors.txt"
	errorFile, err := os.OpenFile(errorFilePath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error creating error file:", err)
		return
	}
	defer errorFile.Close()

	errorCountWriter := &ErrorCountWriter{
		Writer:    os.Stdout,
		ErrorFile: errorFile,
	}

	// Create a TeeReader to read from the log file and write to the custom writer
	teeReader := io.TeeReader(logFile, errorCountWriter)

	// Use bufio.Scanner to read lines from the TeeReader
	scanner := bufio.NewScanner(teeReader)

	// Count of 404 errors
	count404 := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains a 404 error
		if strings.Contains(line, "404") {
			count404++
			// Write the 404 error to the error file
			errLine :=  []byte(line + "\n")
			_, _ = errorCountWriter.ErrorFile.Write(errLine)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading log file:", err)
		return
	}

	fmt.Printf("Number of 404 errors: %d\n", count404)
}
