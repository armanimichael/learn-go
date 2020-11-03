package main

import "fmt"

func main() {
	/*
		Arrays in go have a fixed length and they must
		have the same data type.
		Arrays are 0 based (index starts from 0)
	*/

	// Array of 10 integers
	var array [10]int
	// Setting first two elements
	array[0] = 1
	array[1] = 2
	// All other elements will be zero values of the data type
	fmt.Println("Array of integers:", array)
	fmt.Println("Array's last element:", array[9])

	// Explicitly defining an array
	cities := [3]string{"Milan", "Paris", "Dublin"}
	fmt.Println("Some cities:", cities)

	// Automatically setting array's size on initialization
	languages := [...]string{"Italian", "French", "English"}
	fmt.Println("Some languages:", languages)
}
