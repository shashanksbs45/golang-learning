package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	p := Person{Name: "John", Age: 30, Address: "123 Main St"}
	t := reflect.TypeOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		fmt.Printf("Field: %s, JSON Tag: %s\n", field.Name, tag)
	}
}
