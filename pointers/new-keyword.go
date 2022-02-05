package main

import "fmt"

func main() {
	/*
		It's possible to create a pointer to
		a variable yet to be defined
		by using the new keyword

		In this case we're creating a pointer
		to a string type
	*/
	p := new(string)

	// We can than assign a value to it
	// by using deferencing
	*p = "some value"

	fmt.Printf("Value of *p = '%v'", *p)
}
