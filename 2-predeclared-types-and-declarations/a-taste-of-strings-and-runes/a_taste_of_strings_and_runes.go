package main

import "fmt"

func main() {
	var myFirstInitial rune = 'J' // Good - The type name matches the usage
	var myLastInitial int32 = 'B' // Bad - Legal but confusing
	var cumprimento string = "Olá"

	fmt.Printf("\n%x\n", myFirstInitial) // 4a - Valor numérico em hexadecimal (base 16)
	fmt.Printf("%x\n\n", myLastInitial)  // 4b

	fmt.Printf("%c\n", myFirstInitial)
	fmt.Printf("%c\n\n", myLastInitial)

	fmt.Println(len(cumprimento)) // 4
}
