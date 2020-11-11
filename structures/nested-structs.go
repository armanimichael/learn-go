package main

import "fmt"

func main() {
	/*
		It's possible to define assign a struct to another struct field.
		Doing this is similar to embedding but more verbose. (see embedding-structs.go)
	*/
	type person struct {
		name     string
		lastname string
		age      int
	}

	// We added a person field of type `person` inside worker
	type worker struct {
		person    person
		job       string
		workHours int
	}

	// 1 way to instantiate an embedded struct
	jhon := worker{}
	jhon.person.name = "Jhon"
	jhon.person.lastname = "Smith"
	jhon.person.age = 38
	jhon.job = "programmer"
	jhon.workHours = 40

	// Another way to instantiate an embedded struct, this looks the same as embedding
	jane := worker{
		person:    person{name: "Jane", lastname: "Smith", age: 34},
		job:       "programmer",
		workHours: 40,
	}

	fmt.Println(jhon)
	fmt.Println(jane)
}
