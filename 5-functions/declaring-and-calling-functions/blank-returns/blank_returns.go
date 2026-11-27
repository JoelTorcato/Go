package main

import (
	"errors"
	"fmt"
)

func sumTwoIntegers(value1, value2 int) (result int, subtraction int, err error) {
	if value1 == 0 && value2 == 0 {
		err = errors.New("cannot sum two 0")
		return value1, value2, err
	}
	result, subtraction = value1+value2, value1-value2
	return
}

func main() {
	// Never Use These!

	/*
		If you have named return values,
		you can just write return without specifying the values that are returned.
		This returns the last values assigned to the named return values.
	*/
	fmt.Println(sumTwoIntegers(2, 2)) // 4 0 nil

	result, subtraction, err := sumTwoIntegers(5, 5)
	fmt.Println(result, subtraction, err) // 10 0 nil

	result, subtraction, err = sumTwoIntegers(0, 0)
	fmt.Println(result, subtraction, err) // 0 0 cannot sum two 0
}
