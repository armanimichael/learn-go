package main

import "fmt"

/*
	In Go, logical opearators are described through symbols,

	NOT = !
	OR  = ||
	AND = &&

	----------------

	PS: The \t you see in the prints represent a tab (nothing to do with logical operators)
*/
func main() {
	// The NOT (!) operator makes a boolean condition the opposite
	fmt.Println("NOT: \t!false\t\t", !false) // true
	fmt.Println("NOT: \t!true\t\t", !true)   // false

	// The OR (||) operator returns true when any boolean is equal to true
	fmt.Println("OR: \tfalse || false\t", false || false) // false
	fmt.Println("OR: \ttrue  || true\t", true || true)    // true
	fmt.Println("OR: \tfalse || true\t", false || true)   // true

	// The AND (&&) operator returns true when both booleans are equal to true
	fmt.Println("AND: \tfalse && false\t", false && false) // false
	fmt.Println("AND: \ttrue  && true\t", true && true)    // true
	fmt.Println("AND: \tfalse && true\t", false && true)   // false
}
