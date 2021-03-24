package main

import "fmt"

func main() {
	/*
		Sometimes you may want to create a loop that lasts indefinitely,
		you'll still have to leave the loop at a certain point.

		To do so, you can break the loop.
	*/
	fmt.Println("\nBreak")
	i := 0
	for {
		if i == 10 {
			break
		}
		i++
		fmt.Println(i)
	}

	/*
		We can also skip specific loop's interations
		by using the continue keyword

		In this example, we're skipping odd numbers
	*/
	fmt.Println("\nContinue")
	for i = 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}

	/*
		Breaking loops is also possible in nested loops
		by using labels

		To use label, all you have to do is giving it a name
		of your choice and follow it with :

	*/
	fmt.Println("\nNested Loops")
	for i = 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j == 4 {
				goto Freedom
			}
			fmt.Println(i, j)
		}
	}
Freedom:
}
