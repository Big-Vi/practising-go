package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Student represents a student's information.
type Student struct {
	Age      int
	Username string
}

func main() {
	// Example 1: Working with slices and unsafe.Pointer

	// Create a slice of integers.
	b := []int{3, 5, 6}
	fmt.Println("Original Slice:", b)

	// Get an unsafe pointer to the first element of the slice.
	unsafePtr := unsafe.Pointer(&b[0])
	fmt.Println("Unsafe Pointer to Slice:", unsafePtr)

	// Print the type of the unsafe pointer.
	fmt.Println("Type of Unsafe Pointer:", reflect.TypeOf(unsafePtr))

	// Convert the unsafe pointer to another type (int64 in this case).
	convertUnsafePtr := (*int64)(unsafe.Pointer(&b[0]))
	fmt.Println("Converted Unsafe Pointer:", convertUnsafePtr)

	// Print the type of the converted unsafe pointer.
	fmt.Println("Type of Converted Unsafe Pointer:", reflect.TypeOf(convertUnsafePtr))

	// Dereference the converted unsafe pointer to access the value.
	derefConvertedUnsafePtr := *convertUnsafePtr
	fmt.Println("Dereferenced Converted Unsafe Pointer:", derefConvertedUnsafePtr)

	// Example 2: Pointer arithmetic and type conversion

	// Calculate the address of the next element using pointer arithmetic.
	fmt.Println("Size of Element:", unsafe.Sizeof(&b[0]))
	fmt.Println("Current Address:", uintptr(unsafePtr))
	fmt.Println("Next Address:", uintptr(unsafePtr)+unsafe.Sizeof(&b[0]))

	// Convert the address back to an unsafe pointer.
	nextUnsafePtr := unsafe.Pointer(uintptr(unsafePtr) + unsafe.Sizeof(&b[0]))
	fmt.Println("Next Unsafe Pointer:", nextUnsafePtr)

	// Dereference the next unsafe pointer to access the value.
	fmt.Println("Dereferenced Next Unsafe Pointer:", *(*int64)(nextUnsafePtr))

	// Example 3: Type conversion

	// Convert an integer to a float64 using unsafe.
	i := 60 // Integer
	fmt.Println("Type of Integer:", reflect.TypeOf(i))

	// Perform the conversion using unsafe.Pointer and type casting.
	iToFloat := *(*float64)(unsafe.Pointer(&i))
	fmt.Println("Converted to Float64:", iToFloat)
	fmt.Println("Type of Converted Float64:", reflect.TypeOf(iToFloat))

	// Example 4: Using unsafe with a struct

	// Print the size and field offsets of a struct using unsafe.
	fmt.Println("Size of Student Struct:", unsafe.Sizeof(Student{}))
	fmt.Println("Offset of Age Field:", unsafe.Offsetof(Student{}.Age))
	fmt.Println("Offset of Username Field:", unsafe.Offsetof(Student{}.Username))
}
