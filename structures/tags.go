package main

import (
	"fmt"

	// You can access tags by using the reflect package
	"reflect"
)

func main() {
	/*
		Tags are defined inside backticks ``
		They describe some info about each field.
	*/
	type person struct {
		name     string `example:"Steve"`
		lastname string `example:"James"`
		age      int    `min:"0"`
	}

	// Accessing Tags with reflect
	t := reflect.TypeOf(person{})

	field, ok := t.FieldByName("name")
	tag := field.Tag

	if ok {
		fmt.Println(tag)
	} else {
		fmt.Println("This field name doesn't exist.")
	}
}
