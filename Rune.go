package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	// Create a slice of runes containing a smiley face
	smiley := []rune("🙂")
	fmt.Println(smiley)

	// Print the type of the 'smiley' variable
	fmt.Println(reflect.TypeOf(smiley))

	// Another way to print the type of 'smiley' using Printf
	fmt.Printf("%T", smiley)

	fmt.Println()

	// Check if the smiley face is in the string
	fmt.Println(strings.ContainsRune("🙂", smiley[0]))

	// Define a rune 't' representing the letter 'A'
	t := 'A'
	fmt.Println(byte(t))

	// Define a string 'inputString' containing some Unicode characters
	inputString := "Hello ÀÀ"
	runeSlice := []rune(inputString)

	// Count the number of runes in 'inputString' using utf8.RuneCount
	runeCount := utf8.RuneCount([]byte(inputString))
	fmt.Println(runeCount)

	// Count the number of runes in 'inputString' using utf8.RuneCountInString
	runeCountInString := utf8.RuneCountInString(inputString)
	fmt.Println(runeCountInString)

	fmt.Println("Rune Slice:", runeSlice)

	// Split 'runeInput' into words using custom field splitting function
	runeInput := "Hello world!!! Hi"
	words := strings.FieldsFunc(runeInput, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})
	for _, word := range words {
		fmt.Println(word)
	}

	// Generate a random password string
	allowedCharacters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	password := make([]rune, 36)
	for index := range password {
		password[index] = allowedCharacters[rand.Intn(len(allowedCharacters))]
	}
	fmt.Println(string(password))
}
