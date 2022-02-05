package main

import "fmt"

func main() {
	/*
		When you need to access the value of the variable
		starting from a pointer to that variable, you can
		use deferencing.
	*/
	a := 1
	aPointer := &a

	/*
		In this case, using the * symbol before a pointer
		literally means getting the value of the variable
		in that specific space in memory
		`a` == `aValue`
	*/
	aValue := *aPointer
	fmt.Printf("a value = %v", aValue)
}
