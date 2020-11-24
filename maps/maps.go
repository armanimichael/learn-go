package main

import "fmt"

func main() {
	/*
		Maps are a data type that maps a key to a value,
		they provide a fast and efficient way to access data.
		Duplicated keys are not allowed.
		You can use any type with maps, even structs.

		Underneath, we're defining a map of strings with a integer key
	*/
	var people = map[int]string{
		0:   "Jimmy",
		1:   "John",
		123: "Paul",
	}

	fmt.Println("People:", people)

	// Accessing single keys
	fmt.Println("Person (key = 0)", people[0])
	fmt.Println("Person (key = 1)", people[1])
	fmt.Println("Person (key = 123)", people[123])

	// You can also create maps with make, this will create an empty map
	colors := make(map[string]string)

	// Adding a value to the map
	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	fmt.Println("Colors:", colors)

	// Accessing values from a key
	white := colors["white"]
	fmt.Println("The hex value for white is", white)
}
