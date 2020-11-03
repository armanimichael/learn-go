package main

import "fmt"

func main() {
	/*
		Slices are similar to arrays but their size is not fixed.
		Their capacity can change dinamically.
		Note that slices pass by reference (arrays don't).
	*/

	// Intializing a slice to some values
	slice := []int{1, 2, 3}

	/*
		In here we're declaring a slice
		containing a reference to the first element of the initial slice.
		(See index-filtering.go)
	*/
	slice2 := slice[:1]
	slice2[0] = 0

	// Both slice and slice2 changed as their passed by reference
	fmt.Printf("Slice: %v = Slice2: %v\n", slice, slice2)

	// Declaring an uninitialized
	slice3 := []int{}
	fmt.Println("Slice3:", slice3)

	/*
		Appending elements
		1st arg = Slice
		2nd to nth args = elements to append
	*/
	slice3 = append(slice3, 161, 179, 7, 75, 280)

	// Appending another slices's elements
	slice3 = append(slice3, slice...)

	fmt.Println("Slice3 appends:", slice3)

	// Removing first element
	slice3 = slice3[1:]

	// Removing last element
	slice3 = slice3[:len(slice3)-1]

	fmt.Println(slice3)
	// Removing any item -> in this case the 3rd
	// At this point the slice = [179 7 75 280 0 2]
	slice3 = append(slice3[:3], slice3[4:]...)

	fmt.Println("Slice3 updated:", slice3)

	/*
		The make function let's us create an uninitialized slice with set capacity and length
		1st arg = type
		2nd arg = length
		3rd arg = capacity
	*/
	slice4 := make([]int, 3, 10)
	fmt.Println("Slice4:", slice4)

	// Updating first value
	slice4[0] = 12
	fmt.Println("Slice4 updated:", slice4)
}
