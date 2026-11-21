package main

import "fmt"

func main() {
	words := []string{"hi", "salutations", "hello"}

	for _, word := range words {
		switch wordLen := len(word); { // black switches
		case wordLen < 5:
			fmt.Println(word, "is a short word")
		case wordLen > 10:
			fmt.Println(word, "is a long word")
		default:
			fmt.Println(word, "is exactly the right length")
		}
	}

	/*
		Just as with a regular switch statement, you can optionally include a short variable
		declaration as part of your blank switch. But unlike a regular switch, you can write
		logical tests for your cases. Blank switches are pretty cool, but don't overdo them.
		If you find that you have written a blank switch where all your cases are equality
		comparisons against the same variable you should replace it with an expression switch statement.
	*/
}
