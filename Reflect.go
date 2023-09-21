package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Student represents a student structure with JSON tags.
type Student struct {
	Username string `json:"username"`
	Age      int64  `json:"age"`
}

func main() {
	x := 68
	var y float64
	y = 68

	// Use reflect.DeepEqual to compare values, including their types/structures.
	fmt.Println("Deep equal var", reflect.DeepEqual(x, y)) // False

	// Comparing two maps
	myFruitMap := map[string]int{
		"apple":  20,
		"orange": 40,
	}
	myVeggieMap := map[string]int{
		"carrot":  20,
		"brinjal": 40,
	}
	fmt.Println("Deep equal map", reflect.DeepEqual(myFruitMap, myVeggieMap))

	// Using reflection to inspect variables
	objValuePtr := reflect.ValueOf(&x)
	objType := reflect.TypeOf(&x)
	fmt.Println(objType)
	fmt.Println(objValuePtr)
	fmt.Println(objType.Kind())

	objValue := objValuePtr.Elem()
	fmt.Println(objValue)
	fmt.Println(objValue.CanSet())
	objValue.Set(reflect.ValueOf(40))
	fmt.Println(objValue)

	s := Student{
		Username: "Vicky",
		Age:      34,
	}
	fmt.Println(reflect.ValueOf(s))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(s).Kind())

	// Encode a struct into JSON
	fmt.Println(JSONEncode(s))

	bytes, _ := JSONEncode(s)
	fmt.Println(string(bytes))
}

// JSONEncode encodes a struct into JSON format.
func JSONEncode(v interface{}) ([]byte, error) {
	objType := reflect.TypeOf(v)
	objValue := reflect.ValueOf(v)
	objKind := objType.Kind()
	buf := bytes.Buffer{}

	if objKind != reflect.Struct {
		return buf.Bytes(), fmt.Errorf("value of kind %s is not supported", objKind)
	}

	var output []string
	buf.WriteString("{")
	for i := 0; i < objType.NumField(); i++ {
		fmt.Println(objType.Field(i).Tag.Get("json"))
		switch objValue.Field(i).Kind() {
		case reflect.String:
			output = append(output, string(objType.Field(i).Tag.Get("json"))+":"+objValue.Field(i).Interface().(string))
		case reflect.Int64:
			output = append(output, string(objType.Field(i).Tag.Get("json"))+":"+strconv.FormatInt(objValue.Field(i).Interface().(int64), 10))
		}
	}

	buf.WriteString(strings.Join(output, ","))
	buf.WriteString("}")
	return buf.Bytes(), nil
}
