package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(num, denom int) (int, int, error) { // By convention, the error is always the last (or only) value returned from a function
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil // you must return all of them, separated by commas.
}

func main() {
	result, remainder, err := divAndRemainder(6, 2)

	// If the function complete successfully, you return nil for the error's value.
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder, err) // 3 0 nil
	fmt.Println(remainder)              // 0
	fmt.Println(divAndRemainder(10, 5)) // 2 0 nil
	fmt.Println(divAndRemainder(10, 0)) // 0 0 cannot divide by zero

	result, remainder, err = divAndRemainder(100, 10)

	fmt.Println(result, remainder) // 10 0
}
