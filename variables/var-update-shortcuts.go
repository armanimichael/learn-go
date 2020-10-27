package main

import "fmt"

func main() {
	/*
	  It's possible to update a declared variable by
	  computing an arithmetic operation in terms of
	  the changed variable itself.
	*/
	daysOfWork := 12
	daysOfWork += 4

	fmt.Println(daysOfWork) // 16

	/*
		Similarly, you can use:
		-= sub
		/= div
		*= mult
		%= module (or reminder)
	*/
}
