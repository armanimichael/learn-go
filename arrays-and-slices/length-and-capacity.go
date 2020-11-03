package main

import "fmt"

func main() {
	/*
		You can check array and slices's length and capacity
		by using the len and cap built-in methods
	*/

	// Explicitally setting setting a capacity of 10, with a length of 3 (see slices.go)
	slice := make([]int, 3, 10)

	// The length represents the number of elements initialized
	fmt.Println("Slice's length", len(slice))

	// The capacity represents the size of the underlying array
	fmt.Printf("Slice's capacity %v", cap(slice))
}
