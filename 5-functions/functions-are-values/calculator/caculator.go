package main

import (
	"fmt"
	"strconv"
)

func add(i, j int) int {
	return i + j
}

func sub(i, j int) int {
	return i - j
}

func mul(i, j int) int {
	return i * j
}

func div(i, j int) int {
	return i / j
}

func rem(i, j int) int {
	return i % j
}

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
	"%": rem,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"10", "%", "200"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression: ", len(expression))
			continue
		}
		p1, err := strconv.Atoi(expression[0]) // You're using the strconv.Atoi function in the standard library to convert a string to an int.
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]

		if !ok { // ok != true
			fmt.Println("unsupported operator: ", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(result)
	}

	/*
		You use op as the key to the opMap map, and assign the value associated with the key
		to the variable opFunc. The type of opFunc is func(int, int) int.
		If there wasn't a function in the map associated with the provided key, you print an error
		message and skip the rest of the loop.
	*/
}
