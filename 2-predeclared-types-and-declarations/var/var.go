package main

import "fmt"

func main() {
	x := 10
	x, y := 30, "Hello" // := is not legal outside of functions

	fmt.Printf("\n\n%v\n", x) // %d - Indicar o formato explicitamente
	fmt.Printf("%v\n\n", y)   // %s

	/*
		The most common declaration style within functions is :=.
		Outside of a function, use declaration lists on the rare occasions
		when you're declaring multiple package-lever variables.
	*/
}
