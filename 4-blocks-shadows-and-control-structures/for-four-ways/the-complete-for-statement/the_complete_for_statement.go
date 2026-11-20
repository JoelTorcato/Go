package main

import "fmt"

func main() {
	for i := 0; i < 10; i = i + 2 { // The for statement does not use parentheses around its parts.
		// You must use := to initialize the variables; var is not legal here.
		fmt.Println(i) // 0 2 4 6 8
	}

	/*
	   Go allows you to leave out one or more of the three parts of the for statement.
	   Most commonly, you'll either leave off the initialization if it is based
	   on a value calculated before the loop.
	*/

	a := 0

	for ; a < 10; a++ {
		fmt.Println(a) // 1 2 3 4 5 6 7 8 9
	}

	/*
		Or you'll leave off the increment because you have a more
		complicated increment rule inside the loop
	*/

	for b := 0; b < 10; {
		fmt.Println(a) //

		if b%2 == 0 {
			b++
		} else {
			b += 2
		}
	}
}
