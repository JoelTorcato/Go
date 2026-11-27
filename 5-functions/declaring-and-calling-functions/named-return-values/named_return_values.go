package main

import (
	"errors"
	"fmt"
)

func divAndRemainder(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func main() {
	remainder, result, err := divAndRemainder(100, 10)

	fmt.Println(result, remainder, err) // 0 10 nil
	fmt.Println(result)                 // 0

	fmt.Println(divAndRemainder(10, 2)) // 5 0 nil

	result, remainder, err = divAndRemainder(50, 2)

	fmt.Println(result, remainder, err) // 25 0 nil
	fmt.Println(result)                 // 25
	fmt.Println(remainder)              // 0

	/*
		If you want to name only some of the return values,
		you can do so by using _ as the name for any return values
		you want to remain nameless.
	*/
}
