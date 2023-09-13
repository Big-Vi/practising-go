package main

import (
	"fmt"
	"maps"
)

func main() {
	// Creating and manipulating a map.
	menu := make(map[string]int)
	menu["chicken"] = 60
	menu["fish"] = 100
	menu["lamb"] = 90

	// Initializing map like below creates nil value. When assigning values, it would throw an error.
	// var menu map[string]int

	// Printing the map.
	fmt.Println("Menu:", menu)

	// Deleting an item from the map.
	delete(menu, "fish")
	fmt.Println("Menu after deleting 'fish':", menu)

	// Getting a value from the map.
	fmt.Println("Price of chicken:", menu["chicken"])

	// Finding the length of the map.
	fmt.Println("Number of items in the menu:", len(menu))

	// Checking if a key exists in the map.
	_, keyExists := menu["fish"]
	fmt.Println("Does 'fish' exist in the menu?", keyExists)

	// Iterating over the map.
	fmt.Println("Menu items and their prices:")
	for key, value := range menu {
		fmt.Println("Item:", key, "Price:", value)
	}

	// Creating a map with a struct key type.
	type Menu struct {
		name, category string
	}
	SpecialMenu := map[Menu]float64{
		{name: "Chicken biryani", category: "Lunch"}: 80.99,
		{name: "Lamb biryani", category: "Lunch"}:    100.99,
	}
	fmt.Println("Price of Chicken biryani:", SpecialMenu[Menu{name: "Chicken biryani", category: "Lunch"}])

	// Using functions from the 'maps' package.
	clonedMap := maps.Clone(menu)
	fmt.Println("Cloned map:", clonedMap)

	anotherMenu := make(map[string]int)
	maps.Copy(anotherMenu, menu)
	fmt.Println("Copied menu:", anotherMenu)

	maps.DeleteFunc(menu, func(k string, v int) bool {
		return k == "chicken"
	})
	fmt.Println("Menu after deleting 'chicken':", menu)

	isEqual := maps.Equal(menu, anotherMenu)
	fmt.Println("Is 'menu' equal to 'anotherMenu'?", isEqual)

	isEq := maps.EqualFunc(menu, anotherMenu, func(k int, v int) bool {
		return k == v
	})
	fmt.Println("Is 'menu' equal to 'anotherMenu' using a custom equality function?", isEq)
}
