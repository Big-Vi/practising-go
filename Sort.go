package main

import (
	"fmt"
	"sort"
)

func main() {
	// Sort an integer slice in ascending order
	s := []int{5, 6, 3}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println("Sorted slice s in ascending order:", s)

	// Sort a string slice in lexicographical (alphabetical) order
	s2 := []string{"a", "c", "b"}
	sort.Strings(s2)
	fmt.Println("Sorted slice s2 in ascending order:", s2)

	// Reverse the sorted string slice to get it in descending order
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] > s2[j]
	})
	fmt.Println("Sorted slice s2 in descending order:", s2)

	// Use sort.Search to find the index of a specific element 400 which performs binary search
	s3 := []int{67, 98, 300, 50, 4, 79, 400, 50}
	index := sort.Search(len(s3), func(i int) bool {
		return s3[i] >= 400
	})
	fmt.Println("Index of 400:", index)
}
