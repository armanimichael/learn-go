package main

import "fmt"

func main() {
	/*
		As Go doesn't have inheritance
		we can use a concept called Composition
	*/
	type person struct {
		name     string
		lastname string
		age      int
	}

	// We added `person` inside worker, the worker struct has now `person` fields
	type worker struct {
		person
		job       string
		workHours int
	}

	// 1 way to instantiate an embedded struct
	john := worker{}
	john.name = "John"
	john.lastname = "Smith"
	john.age = 38
	john.job = "programmer"
	john.workHours = 40

	// Another way to instantiate an embedded struct
	jane := worker{
		person:    person{name: "Jane", lastname: "Smith", age: 34},
		job:       "programmer",
		workHours: 40,
	}

	fmt.Println(john)
	fmt.Println(jane)
}
