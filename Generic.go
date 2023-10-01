package main

import (
	"fmt"
	"slices"
	"golang.org/x/exp/constraints"
)

// Number is an interface representing int64 or float64 types.
type Number interface {
	int64 | float64
}

// NumberFromConstraint is an interface representing constraints.Integer or constraints.Float.
type NumberFromConstraint interface {
	constraints.Integer | constraints.Float
}

// Dinner represents a meal with a name and price.
type Dinner struct {
	Name  string
	Price int
}

// Type Parameters
func GetString[T any](t T) string {
	return fmt.Sprintf("%s - %v", "MyFood", t)
}

// Type Constraints
func GetMenuDetails[K comparable, V int|string](t map[K]V) []V {
	output := []V{}
	for _, v := range t {
		output = append(output, v)
	}
	return output
}

func GetTotal[K comparable, V NumberFromConstraint](m map[K]V) V {
	var output V
	for _, v := range m {
		output += v
	}
	return output
}

// Type Sets
type Car struct {
	Brand string
	Model string
}

func (c Car) GetWheels() int {
	return 4
}

type Bike struct {
	Brand string
	Model string
}

func (b Bike) GetWheels() int {
	return 2
}

type Vehicle interface {
	Car | Bike
	GetWheels() int
}

func GetVehicle[T Vehicle](t T) string {
	return fmt.Sprintf("%s - %v", t, t.GetWheels())
}

// Constraints package
type Filter constraints.Ordered

func FilterArray[T Filter](a []T) []T {
	var out []T
	slices.Sort(a)
	for _, v := range a {
		if !slices.Contains(out, v) {
			out = append(out, v)
		}
	}
	return out
}

// Tilde
type CustomType int64

func printValue[T ~int64](t T) {
	fmt.Println("Value: ", t)
}

type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
    *s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
    if len(*s) == 0 {
        panic("stack is empty")
    }
    value := (*s)[len(*s)-1]
    *s = (*s)[:len(*s)-1]
    return value
}

func main() {
	dinner := Dinner{Name: "Fish curry", Price: 40}
	Value1 := GetString(dinner)
	fmt.Println(Value1)
	Value2 := GetString("Tasty")
	fmt.Println(Value2)

	// Type constraint
	m := map[int]string{1: "Fish", 2: "Chicken", 3: "Lamb"}
	Value3 := GetMenuDetails(m)
	fmt.Println(Value3)
	m1 := map[string]int{"Fish": 100, "Chicken": 50, "Lamb": 40}
	Value4 := GetMenuDetails(m1)
	fmt.Println(Value4)

	m3 := map[string]int64{"Fish": 100, "Chicken": 50, "Lamb": 40}
	fmt.Println(GetTotal(m3))

	m4 := map[string]float64{"Fish": 100.50, "Chicken": 50.45, "Lamb": 40}
	fmt.Println(GetTotal(m4))

	// Type set
	v := Car{Brand: "Nissan", Model: "2012"}
	fmt.Println(GetVehicle(v))
	v1 := Bike{Brand: "Hero"}
	fmt.Println(GetVehicle(v1))

	// Constraints Package
	array := []int{1, 3, 3, 4, 6}
	array2 := []int32{1, 3, 3, 4, 4, 7, 6}
	fmt.Println(FilterArray(array))
	fmt.Println(FilterArray(array2))

	// Tilde
	var myValue CustomType
	myValue = 4000
	printValue(myValue)

	// Generic data structure
    stack := Stack[int]{}
    stack.Push(10)
    stack.Push(20)
    fmt.Println(stack.Pop()) // Outputs: 20
	fmt.Println(stack)

	stack2 := Stack[string]{}
    stack2.Push("string1")
    stack2.Push("string2")
    fmt.Println(stack2.Pop()) // Outputs: string2
	fmt.Println(stack2)
}
