package main

import "fmt"

func main() {
	/*
		Loops are quite useful when it
		comes to accessing collections, such as the slice below
	*/
	fmt.Println("\nFor Loop")
	list := []string{"a", "b", "c"}
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}

	/*
		You can also access collections
		by using the range keyword
	*/
	fmt.Println("\nFor Range Loop")
	for index, item := range list {
		fmt.Println(index, item)
	}
}
