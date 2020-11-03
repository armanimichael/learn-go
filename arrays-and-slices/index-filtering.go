package main

import "fmt"

func main() {
	/*
		You can filter arrays to return a range of values
	*/
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Returns all elements after the third one (not included)
	fmt.Println("Same as using only the array name:", a[:])

	// Returns all elements after the third one (included)
	fmt.Println("After 3rd element\t\t:", a[3:])

	// Returns all elements before the third one (excluded)
	fmt.Println("Before 3rd element\t\t:", a[:3])

	// Returns all elements between third and eighth ones (third included, eighth excluded)
	fmt.Println("Between 3rd and 8th element\t:", a[3:8])
}
