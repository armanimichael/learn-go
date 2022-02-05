package main

import "fmt"

func main() {
	// It is possible to get a pointer of a pointer
	// since pointer still take space in memory
	a := 1
	aPointer := &a
	aPointerPointer := &aPointer // I know this naming is terrible, no other ideas...

	fmt.Printf(" a address = %v\n", aPointer)
	fmt.Printf("&a address = %v\n", aPointerPointer)
}
