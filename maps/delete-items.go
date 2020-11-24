package main

import "fmt"

func main() {
	/*
		Deleting a value from a map is possible.
		To do so, you can use the delete function
	*/
	var people = map[int]string{
		0:   "Jimmy",
		1:   "John",
		123: "Paul",
	}

	// delete accepts a map and an index
	delete(people, 123)

	fmt.Println(people)
}
