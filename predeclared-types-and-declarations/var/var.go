package main

import "fmt"

func main() {
	x := 10
	x, y := 30, "Hello"

	fmt.Printf("\n\n%v\n", x) // %d - Indicar o formato explicitamente
	fmt.Printf("%v\n\n", y)   // %s
}
