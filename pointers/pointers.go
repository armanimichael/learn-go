package main

import "fmt"

func main() {
	/*
		Pointers allow you to create a referece
		to a specific type address in memory.
		You can have a pointer to anything
		from simple types to function, interfaces and structs
	*/
	a := 1

	// The & symbol means 'address of'.
	// In this case we're accessing the (memory) address of variable a
	aPointer := &a

	fmt.Printf("Value of a: %v\n", a)
	fmt.Printf("Address of a: %v\n", aPointer)
}
