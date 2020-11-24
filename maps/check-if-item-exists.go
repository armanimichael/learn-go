package main

import "fmt"

func main() {
	/*
		When you access a map's value with a key, you'll always get a value.
		Accessing a non-existing key will return a zero value.
		To check if that key exists, you can use the ok variable (can be named whatever, this is just a convention)
	*/
	var people = map[int]string{
		0:   "Jimmy",
		1:   "John",
		123: "Paul",
	}

	// person is a string, while ok is a boolean
	_, ok := people[2]

	// This will print false (because there's no value at key 2)
	fmt.Println(ok)
}
