/*
	The Universe Block

	fmt.Println(true) // true
	true := 10
	fmt.Println(true) // 10

	You must be very careful to never redefine any identifiers in the universe block.
*/

package main

import "fmt" // Bloco do ficheiro - O import disponibiliza o nome fmt neste ficheiro

func main() { // Bloco do pacote - A função main é declarada neste bloco.
	x := 10 // Bloco do corpo da função

	if x > 5 {
		fmt.Println(x) // 10
		x := 5         // Bloco interno - corpo deste if.
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10

	a := 10

	if a > 5 {
		a, b := 5, 20     // := reuses only variables that are declared in the current block
		fmt.Println(a, b) // 5 20
	}

	fmt.Println(a) // 10
	/*
		fmt := "oops"
		fmt.Println(fmt)
		fmt.Println undefined (type string has no field or method Println)
		Notice that the problem isn't that your named your variable fmt;
		it's that you tried to access something that the local variable fmt didn't have.
	*/
}
