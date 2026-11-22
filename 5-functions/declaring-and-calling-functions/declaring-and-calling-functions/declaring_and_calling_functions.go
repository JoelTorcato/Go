package main

import "fmt"

func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

func main() {
	/*
		A function declaration has four parts:
		- The keyword func
		- The name of the function
		- The input parameters (are listed in parentheses, separated by commas, with the parameter name first and the type second).
		- The return type
	*/
	result := div(10, 5)
	fmt.Println(result) // 2

	// fmt.Println(div(10, 5))
}
