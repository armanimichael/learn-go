package main

import "fmt"

func main() {
	a, b, c := "Go", "formatted", "text!"

	/*
		The Println method accepts multiple arguments
		they will display one after another and
		automatically separate with black space.
	*/
	fmt.Println(a, b, c)
	fmt.Println("This", "also", "works")

	/*
		Similarly, you can use Print to concatenate arguments
		without adding black spaces and without any escape
		character at the end
	*/
	fmt.Print(a, b, c)
	fmt.Print("This", "also", "works")
	fmt.Println()

	/*
		You can also create a variable containing the formatted text by using the Sprintln or Sprint methods
	*/
	sentence := fmt.Sprintln(a, b, c)
	fmt.Print(sentence)

	spacelessSentence := fmt.Sprint(a, b, c)
	fmt.Println(spacelessSentence)

	/*
		Another option is to use 'verbs'.
		Verbs are special characters you can replace with variables.
		To do so, you have to use the Printf method.
			%v reppresents the value of the variable
			%T reppresents the type of the variable
			%#v reppresents the name of the variable
			%f reppresents a floating point
		You can read about the others on Go's docs at https://golang.org/pkg/fmt/
		Also note that Printf doesn't create a new line,so we're adding \n at the end to do it manually
	*/
	d := "verbs"
	fmt.Printf("We're now using %v. Cool, isn't it?\n", d)
	fmt.Printf("The variable %#v is a %T\n", d, d)

	/*
		By using %f, we can specify the number of decimals
		For example, f.2f means two decimals
	*/
	num := 2.4
	fmt.Printf("Unformatted float %f\n", num)
	fmt.Printf("Formatted float %.2f\n", num)

	/*
		You can also use Sprintf
	*/
	text := fmt.Sprintf("%v formatting is now inside a variable", d)
	fmt.Println(text)
}
