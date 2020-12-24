package main

import "fmt"

/*
	'if' is a type of conditional,
	They evaluate boolean expressions => note that we don't use parenthesis
*/
func main() {
	a, b := 1, 2

	// Simple if example
	if a != b {
		fmt.Println("A and B are different")
	}

	// if-else
	if a != b {
		fmt.Println("A and B are different")
	} else {
		fmt.Println("A and B are equal")
	}

	// if-else-else if
	if a != b {
		fmt.Println("These numbers are different")
	} else if a > b {
		fmt.Println("A is greater than B")
	} else if a < b {
		fmt.Println("A is less than B")
	} else {
		fmt.Println("A and B are equal")
	}
}
