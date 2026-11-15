package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d"}
	y := x[1:3]

	fmt.Println(x, cap(x)) // [a b c d] 4
	fmt.Println(y, cap(y)) // [b c] 3 (cap(b) = cap(a) - início (1))

	y = append(y, "z")
	fmt.Println("x:", x) // [a b z d] - Reescreve o valor c
	fmt.Println("y:", y) // [b c z]
}
