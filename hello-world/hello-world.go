/*
Go works by splitting into packages,
you can name your package any name you wish.
The main package is an entry point for the program,
it is the only package that will generate an executable file.

You can build your project using:
	go build hello-world.go
This will generate a hello-world executable file (for example .exe in Windows)

If you want to execute your program right away, without creating an executable,
you can use:
	go run hello-world.go
*/
package main

/*
The import statement lets you import packages.
In particular, the fmt package implements formatted I/O
*/
import "fmt"

/*
You can create function using the func keywork.
The main function will automatically be called once you start the program
*/
func main() {
	// This method from fmt will print a string and escape on a new line
	fmt.Println("Hello World")
}
