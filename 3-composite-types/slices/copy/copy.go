package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5} // Slice permite trabalhar com uma parte de um array e pode mudar de comprimento
	d := [4]int{6, 7, 8, 9}   // Array comprimento fixo
	y := make([]int, len(x))
	a := make([]int, 2)
	b := make([]int, 3)
	e := make([]int, 2)

	copy(e, d[:])  // copy(destino, origem), copy([1 2 3], [2 3 4 5])
	fmt.Println(e) // [6 7], length - 2
	copy(d[:], x)
	fmt.Println(d) // [1 2 3 4], length - 4

	lengthY := copy(y, x)
	lengthA := copy(a, x)
	lengthB := copy(b, x[1:4])
	c := copy(x[:3], x[1:]) // O destino tem apenas 3 posições (length 3), portante recebe apenas os primeiros 3 valores da origem [2 3 4]

	fmt.Println(y, lengthY)   // [1 2 3 4 5] 5
	fmt.Println(a, lengthA)   // [1 2] 2
	fmt.Println(b, lengthB)   // [2 3 4] 3
	fmt.Println(x, c, len(x)) // [2 3 4 4 5] 3 5
}
