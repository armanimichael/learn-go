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
	jhon := worker{}
	jhon.name = "Jhon"
	jhon.lastname = "Smith"
	jhon.age = 38
	jhon.job = "programmer"
	jhon.workHours = 40

	// Another way to instantiate an embedded struct
	jane := worker{
		person:    person{name: "Jane", lastname: "Smith", age: 34},
		job:       "programmer",
		workHours: 40,
	}

	fmt.Println(jhon)
	fmt.Println(jane)
}
