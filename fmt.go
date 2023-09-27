package main

import (
	"fmt"
)

type Student struct {
	Username string
}

func main() {
	age := 30
	name := "Vignesh"

	// Print age using the %v format verb
	fmt.Printf("Age: %v\n", age)

	student := Student{Username: "Vignesh"}
	// Print Student struct with Go-syntax representation
	fmt.Printf("Student: %#v\n", student)

	isTrue := true
	// Print a boolean value
	fmt.Printf("Is True? %t\n", isTrue)

	// Print the type of isTrue variable
	fmt.Printf("Type of isTrue: %T\n", isTrue)

	// Print age with padding of zeros to the left (minimum width of 5 characters)
	fmt.Printf("Padded Age (0): %05d\n", age)

	// Print age with padding of spaces to the left (minimum width of 5 characters)
	fmt.Printf("Padded Age (space): % 5d\n", age)

	// Print name and age
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Store formatted string in a variable
	store := fmt.Sprintf("Formatted: %s, %d", name, age)
	fmt.Println(store)

	// Print a floating-point number with 2 decimal places
	fmt.Printf("Float: %.2f\n", 3.14444)

	// Print name with a minimum width of 10 characters
	fmt.Printf("Padded Name: %10s\n", name)

	char := 65
	// Print age in binary representation
	fmt.Printf("Age in Binary: %b\n", age)

	// Print the character represented by the Unicode code point 65 (A)
	fmt.Printf("Character: %c\n", char)

	// Print the Unicode code point of the character represented by 65 (A)
	fmt.Printf("Unicode Code Point: %U\n", char)

	// Argument indexes
	one := 1
	two := 2
	three := 3
	// Print variables with specified argument indexes
	fmt.Printf("Argument Indexes: %[1]d %[3]d %[2]d %[2]d\n", one, two, three)

	// Print name with a dynamic width of 10 characters
	fmt.Printf("Dynamic Width: %*s\n", 10, name)

	// Print the first 3 characters of the name
	fmt.Printf("First 3 Characters: %.*s\n", 3, name)
}
