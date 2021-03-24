package main

import "fmt"

func main() {
	/*
		1st statement: initializer
		2nd statement : boolean result statement
		3rd statement: incrementer
	*/
	fmt.Println("\nBasic Loop")
	for i := 0; i < 10; i++ {
		fmt.Println(1)
	}

	/*
		Using multiple initialized values
	*/
	fmt.Println("\nTwo variables")
	for i, j := 0, 10; i < 10; i, j = i+1, j-1 {
		fmt.Println(i, j)
	}

	/*
		You can actually define the iterator variable outside
		of the for scope, leaving the first statement empty
		MIND THE SEMICOLON!
	*/
	fmt.Println("\nOuter scope variable")
	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	/*
		You can also remove the third statement.
		In this example we're increasing i inside the loop (otherwise we'd get an infinite loop)
		MIND THE SEMICOLON!
	*/
	fmt.Println("\nIncrementer moved inside the loop")
	for i = 0; i < 10; {
		fmt.Println(i)
		i++
	}

	/*
		In Go the while keyword doesn't exist.
		A while loop is just a special for loop.

		You can have the  same result others languages achieve with while
		by removing the first and third statements
		NO NEED TO WRITE SEMICOLONS
	*/
	fmt.Println("\nWhile loop")
	i = 0
	for i < 10 {
		fmt.Println(i)
	}
}
