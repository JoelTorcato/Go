package main

import "fmt"

func main() {
	for i := 1; i < 15; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	/*
		Favor blank switch statements over if/else chains when you
		have multiple related cases. Using a switch makes the comparisons
		more visible and reinforces that they are a related set of concerns.
	*/
}
