package main

import "fmt"

func main() {
	/*
		Fetching User's inputs is mandatory for interactivity,
		you can use the Scan method to do it.
		The & symbol indicates a pointer (memory address), if your familiar wih C
		you already know what I'm talking about, otherwise,
		I will talk about pointers in the future
	*/
	var input string

	/*
		By running this snippet, you'll be able to write
		something and see it printed after you press enter.
	*/
	fmt.Scan(&input)
	fmt.Println(input)

	/*
		HOWEVER: This only works with the first word,
		if you wish to read multiple words you can
		create multiple variables
	*/
	var firstWord, secondWord string
	fmt.Scan(&firstWord, &secondWord)
	fmt.Println(firstWord, secondWord)

	/*
		You can perform more complex scans using
		Scanf, it will require the use of verbs
		to predermine the form of the wanted input.
		In this case, it accetps a string and an integer => %s %d
	*/
	var firstPart string
	var secondPart int
	fmt.Scanf("%s %d", &firstPart, &secondPart)
	fmt.Println(firstPart, secondPart)
}
