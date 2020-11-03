package main

import "fmt"

func main() {
	/*
		Arrays in Go are NOT passed by reference.
		To clone an array, all you have to do is
		assigning them to another variable
	*/
	a := [...]int{1, 2, 3}
	b := a

	b[0] = 0

	// b changes while a doesn't
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	/*
		If you want to create a reference,
		you must use a pointer (referencing to the original array)
	*/
	c := &a
	c[0] = 0

	// both a and c changed, because c is a reference to a (not an array)
	fmt.Println("a:", a)
	fmt.Println("c:", c)
}
