package main

import "fmt"

func main() {
	// Confusing code

	for i := 1; i <= 15; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz") // Replace multiples of both 3 and 5 (multiple of 15) with "FizzBuzz"
			} else {
				fmt.Println("Fizz") // Replace multiples of 3 with "Fizz"
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz") // Replace multiples of 5 with "Buzz"
		} else {
			fmt.Println(i)
		}
	}

	// Using continue to make code clearer

	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	/*
		As you can see, replacing chains of if/else statements with
		a series of if statements that use continue makes the conditions line up.
		This improves the layout of your conditions, which means your code
		is easier to read and understand.
	*/
}
