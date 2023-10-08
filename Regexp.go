package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Regular expression to match strings in the format: word[numeric]
	pattern := regexp.MustCompile(`^[a-zA-Z]+\[[0-9]+\]$`)

	// Check if a string matches the pattern
	fmt.Println(pattern.MatchString("adam[4555]"))          // true
	fmt.Println(pattern.Match([]byte("adam[4555]")))        // true

	// Replace occurrences of 'x' or 'y' with 'z'
	pattern2 := regexp.MustCompile("x|y")
	fmt.Println(pattern2.ReplaceAllString("xyx", "z"))       // zzz

	// Check if a string matches the pattern (alternative method)
	match, _ := regexp.MatchString(`^[a-zA-Z]+\[[0-9]+\]$`, "fun[99]")
	fmt.Println(match)                                      // true

	// Regular expression to find and capture SSN-like patterns
	pattern3 := regexp.MustCompile(`(\d{3})-(\d{2})-(\d{4})`)

	// Find and capture submatches in a string
	submatches := pattern3.FindStringSubmatch("123-45-6789")
	fmt.Println(submatches)                                 // [123-45-6789 123 45 6789]

	// Find the start and end indices of submatches in a string
	indexSubmatch := pattern3.FindStringSubmatchIndex("My security number 134-56-6768")
	fmt.Println(indexSubmatch)                               // [21 32 21 24 25 32 21 32]

	// Find the start and end indices of the first match in a string
	indexMatch := pattern3.FindStringIndex("My security number 134-56-6768")
	fmt.Println(indexMatch)                                  // [21 32]
}
