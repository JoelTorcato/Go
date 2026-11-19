package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n := rand.Intn(10) // Returns, as an int, a non-negative pseudo-random number in the half-open [0, n).

	if n == 0 { // Don't put parentheses around the condition.
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("Tha's too big", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}
