package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	/*
		Having this special scope is handy.
		It lets you create variables that are available only where they are needed.
		Once the series of if/else statements ends, n is undefined. (out_of_scope.go)
	*/
}
