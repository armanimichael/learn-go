package main

import "fmt"

func main() {
	/*
		As in any programming language, you can store data inside variables
		In Go, we use camelCase to represent variables.
		To create a variable, we can use the var keyword
		For example:
	*/

	var someInteger = 47
	var someFlootingPoint = 24.22
	var name = "Giulio"
	var isGoFun = true

	/*
		We can also create constants, in this case, once their value
		is set, they cannot be changed.
		The only difference is the use of the const keyword
	*/
	const eulerNumber = 2.7182
	const untouchableString = "You can't change me!"

	// We can then print the out in console if we wish to.
	fmt.Println(someInteger)
	fmt.Println(someFlootingPoint)
	fmt.Println(name)
	fmt.Println(isGoFun)
	fmt.Println(eulerNumber)
	fmt.Println(untouchableString)
}
