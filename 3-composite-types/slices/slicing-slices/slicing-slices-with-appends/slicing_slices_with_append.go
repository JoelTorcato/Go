package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[:2]

	fmt.Println(x, cap(x)) // [a b c d] 4
	fmt.Println(y, cap(y)) // [a b] 4

	y = append(y, "z")
	fmt.Println("x:", x) // [a b z d] - Reescreve o valor c
	fmt.Println("y:", y) // [a b z]
}
