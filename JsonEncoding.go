package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Student represents a student with custom time format for CreatedAt field.
type Student struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Age       int         `json:"age"`
	CreatedAt customTime  `json:"createdAt"`
}

// customTime is a custom time type that extends time.Time for custom JSON marshaling and unmarshaling.
type customTime struct {
	time.Time
}

const customTimeFormat = "2006-01-02 15:04:05" // Custom time format

// MarshalJSON customizes the JSON marshaling of customTime.
func (c customTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", c.Time.Format(customTimeFormat))), nil
}

// UnmarshalJSON customizes the JSON unmarshaling of customTime.
func (c *customTime) UnmarshalJSON(data []byte) error {
	var err error
	// Remove double-quotes and parse the time using the custom format
	c.Time, err = time.Parse(customTimeFormat, strings.ReplaceAll(string(data), "\"", ""))
	if err != nil {
		return err
	}
	return nil
}

// Class represents a class with a student and name.
type Class struct {
	Student Student `json:"student"`
	Name    string  `json:"name"`
}

func main() {
	// Create a Class instance with sample data
	class := Class{
		Name: "History",
		Student: Student{
			ID:        "1",
			Name:      "Vicky",
			Age:       33,
			CreatedAt: customTime{time.Now()},
		},
	}

	// Marshal the Class struct to JSON
	bytes, err := json.Marshal(class)
	if err != nil {
		fmt.Println("Error marshaling:", err)
	}
	fmt.Println(string(bytes))

	// Unmarshal the JSON data back to a Class struct
	var c = Class{}
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
	}

	// Access the data from the unmarshaled struct
	fmt.Println("Class Name:", c.Name)
	fmt.Println("Student ID:", c.Student.ID)
	fmt.Println("Student Name:", c.Student.Name)
	fmt.Println("Student Age:", c.Student.Age)
	fmt.Println("Student CreatedAt:", c.Student.CreatedAt.Time)
}
