package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Create a ZIP archive file
	archiveFile := "myarchive.zip"
	if err := createZIPArchive(archiveFile); err != nil {
		log.Fatal("Error creating ZIP archive:", err)
	}

	// Extract files from the ZIP archive
	extractFolder := "extracted_files"
	if err := extractZIPArchive(archiveFile, extractFolder); err != nil {
		log.Fatal("Error extracting ZIP archive:", err)
	}

	fmt.Println("Files extracted successfully.")
}

// createZIPArchive creates a new ZIP archive with one or more files.
func createZIPArchive(zipFileName string) error {
	// Create or open the target ZIP file for writing
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Initialize a ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Files to include in the ZIP archive (you can add more files as needed)
	filesToArchive := []struct {
		Name, Content string
	}{
		{"file1.txt", "This is the content of file1.txt."},
		{"file2.txt", "Content of file2.txt."},
	}

	// Iterate over files and write them to the ZIP archive
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)
		if err != nil {
			return err
		}
		_, err = fileWriter.Write([]byte(file.Content))
		if err != nil {
			return err
		}
	}

	return nil
}

// extractZIPArchive extracts the contents of a ZIP archive into a folder.
func extractZIPArchive(zipFileName, extractFolder string) error {
	// Open the ZIP archive file for reading
	zipFile, err := zip.OpenReader(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create the extraction folder if it doesn't exist
	if err := os.MkdirAll(extractFolder, os.ModePerm); err != nil {
		return err
	}

	// Iterate over the files in the ZIP archive and extract them
	for _, file := range zipFile.File {
		// Open the file inside the ZIP archive
		zipFile, err := file.Open()
		if err != nil {
			return err
		}
		defer zipFile.Close()

		// Create a new file on the local filesystem
		extractFilePath := filepath.Join(extractFolder, file.Name)
		extractFile, err := os.Create(extractFilePath)
		if err != nil {
			return err
		}
		defer extractFile.Close()

		// Copy the contents from the ZIP file to the local file
		_, err = io.Copy(extractFile, zipFile)
		if err != nil {
			return err
		}
	}

	return nil
}
