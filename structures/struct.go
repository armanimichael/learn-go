package main

import "fmt"

func main() {
	/*
		Go is not an Object Oriented language,
		instead of creating classes you can create structures.
		Fields can have any kind of data, even other structs.
		We can also associate method to structs (later on)
	*/
	type person struct {
		name     string
		lastname string
		age      int
	}

	// Creating values from `person` struct
	jimmy := person{
		name:     "Jimmy",
		lastname: "Smith",
		age:      24, // Not the comma after the last field
	}
	fmt.Println(jimmy)
	fmt.Printf("%v %v is %v years old.\n", jimmy.name, jimmy.lastname, jimmy.age)

	// No need to instantiate a struct fields right away, they default to zero values
	johnny := person{}

	//  Accessing and setting single field
	johnny.name = "Johnny"

	fmt.Println(johnny)

	/*
		Structs are values, you can clone
		a struct by simply assigning it to
		a new variable
	*/
	anotherJimmy := jimmy

	// Updating the new var
	anotherJimmy.lastname = "Adams"

	fmt.Println("Original Jimmy:", jimmy)
	fmt.Println("New Jimmy:", anotherJimmy)
}
