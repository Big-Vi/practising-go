package main

import (
	"fmt"
	"slices"
)

func main() {
	// Define an array and assign values
	var array [3]int
	array = [3]int{4, 5, 6}
	fmt.Println("array:", array)

	// Define an array and initialize it
	array2 := [3]int{4, 6, 9}
	fmt.Println("array2:", array2)

	// Set a value in array2
	array2[2] = 60
	fmt.Println("array2:", array2)

	// Define a slice and assign values
	var slice []int
	slice = []int{7, 8, 9}
	fmt.Println("slice:", slice)

	// Use make to create a slice with a specified length and capacity
	slice2 := make([]int, 3, 10)
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))

	// Slicing a slice doesn't create a new slice, it references the underlying slice
	newslice := slice[1:3]
	fmt.Println("newslice:", newslice)
	slice[2] = 10
	fmt.Println("newslice:", newslice)
	fmt.Println("slice:", slice)

	// Convert an array to a slice
	array3 := [3]int{}
	arrayToSlice := array3[:]
	arrayToSlice[1] = 10
	fmt.Println("array to slice:", array3, "cap:", cap(array3))
	fmt.Println("array to slice:", arrayToSlice, "cap:", cap(arrayToSlice))

	// Copy a slice
	slice3 := []int{4, 5, 7}
	slicecopy := make([]int, 3)
	copy(slicecopy, slice3)
	fmt.Println("slice copy:", slicecopy)

	// Slices package

	// Clone - create a new copy of a slice
	sliceclone := slices.Clone(slice3)
	fmt.Println("cloned:", sliceclone)
	sliceclone[2] = 8
	fmt.Println("cloned:", slice3)

	// Grow - increase the capacity of a slice
	fmt.Println("slice3:", slice3, "cap:", cap(slice3))
	slicegrow := slices.Grow(slice3, 5)
	fmt.Println("slice3:", slice3, "cap:", cap(slice3))

	// The slice automatically doubles in capacity when values are appended
	slice3 = append(slice3, 7)
	fmt.Println("slice3:", slice3, "cap:", cap(slice3))
	fmt.Println("slicegrow:", slicegrow, "cap:", cap(slicegrow))

	originalSlice := []int{1, 2, 3, 4, 5}

    // Create a new slice with the same underlying array but length and capacity set to 0
    newSlice := originalSlice[:0:0]

    fmt.Println("Original Slice:", originalSlice)
    fmt.Println("New Slice:", newSlice)
    fmt.Println("Length of New Slice:", len(newSlice))
    fmt.Println("Capacity of New Slice:", cap(newSlice))
}
