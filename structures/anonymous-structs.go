package main

import "fmt"

func main() {
	/*
		Anonymous structs let you create
		a struct on the fly, without a name.
		They won't be accessible anywhere else.
	*/
	person := struct {
		name     string
		lastname string
		age      int
	}{
		name:     "John",
		lastname: "Smith",
		age:      64,
	}

	fmt.Println(person)
}
