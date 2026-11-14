package main

import "fmt"

func main() {
	var myFirstInitial rune = 'J' // Good - The type name matches the usage
	var myLastInitial int32 = 'B' // Bad - Legal but confusing

	fmt.Printf("%x\n", myFirstInitial)
	fmt.Printf("%x\n", myLastInitial)
}
