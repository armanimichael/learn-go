package main

import "fmt"

/*
	'switch' is similar to 'if':
	it's a short way to compare a value agains other values.
	The default keyword works similar to the else inside ifs.
*/
func main() {
	var cmd string
	fmt.Print("Run command: ")
	fmt.Scan(&cmd)

	// Simple switch
	switch cmd {
	case "start":
		fmt.Println("Let's start")
	case "leave":
		fmt.Println("Goodbye")
	default:
		fmt.Println("Command not found")
	}

	// Multiple comparisons with same result
	switch cmd {
	case "start", "begin", "go":
		fmt.Println("Let's start")
	case "leave", "exit", "cancel":
		fmt.Println("Goodbye")
	default:
		fmt.Println("Command not found")
	}

	fmt.Println()
	/*
		You can create more complex switch statements that works
		almost exactly like ifs.
		You can use the fallthrough keyword to also test the other cases.
		In other languages there's a break keyword at the end of each case, in Go, breaks are implicit
	*/
	fmt.Println("Complex switch: ")
	num := 4
	switch {
	case num >= 4:
		fmt.Printf("%v is greater or equal than 4\n", num)
		fallthrough

	case num >= 3:
		fmt.Printf("%v is greater or equal than 3\n", num)
		// We're not using fallthrough here, it won't test the case beneath

	case num >= 2:
		fmt.Printf("%v is greater or equal than 2\n", num)

	}
}
