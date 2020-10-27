/*
Running this snippet will produce an error,
as the declared variables are unused
*/
package main

func main() {
	/*
		When you declare a variable and assign a value straightaway,
		you can avoid assigning a type.
		However, if you want to declare an uninitialized var you MUST declare its type.
		You can read about more types on the Go docs, I'll just stick to basic ones.
	*/
	var integer int32
	var decimal float32
	var stringValue string
	var trueOrFalse bool

	/*
		You can also use int without specifing its size,
		it will automatically set based on the computer architecture 32 or 64 bit
	*/
	var guessTheArchitecture int

	/*
		You can also declare multiple variables of the same types on the same line
	*/
	var a, b, c, d, e, f string

	/*
		For numeric types, there are different types depending on the
		memory they allocate to reppresent a specific number.
		For example:
			int32 allows the reppresentation of 2^32 values (for -2,147,483,648 to 2,147,483,647)
			bool also count as a number with an allocation of 2^1 bits => 0 (false) or 1 (true)
	*/
}
