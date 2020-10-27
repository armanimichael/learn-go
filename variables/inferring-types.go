package main

func main() {
	/*
		As said in `data-types.go` you can create a variable
		and assiging a value to it without specifing its type.
		We call this procedure Type inference.
		It's possible to do that in two ways.
	*/

	// 1st method
	var someInteger = 47
	var someFlootingPoint = 24.22
	var name = "Giulio"
	var isGoFun = true

	// 2nd method
	someInteger2 := 47
	someFlootingPoint2 := 24.22
	name2 := "Giulio"
	isGoFun2 := true

	/*
		You can also infer multiple variables on the same line
	*/
	a, b, c := "Some text", 54, true

	/*
		a = "Some Text"
		b = 54
		c = true
	*/
}
