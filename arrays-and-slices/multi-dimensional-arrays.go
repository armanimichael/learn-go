package main

import "fmt"

func main() {
	/*
		Multi-dimensional arrays are
		basically arrays of arrays.
		The first index indicates rows, the second represents columns
	*/
	table := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// Changing a value on 1rd row and 3rd column
	table[0][2] = 100

	fmt.Println("First table:", table)

	// You can also write this like so
	var table2 [2][3]int
	table2[0] = [3]int{1, 2, 3}
	table2[1] = [3]int{4, 5, 6}

	fmt.Println("Second table:", table2)
}
