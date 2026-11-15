package main

import "fmt"

func main() {
	x := make([]string, 0, 5)         // length 0, cap 5
	x = append(x, "a", "b", "c", "d") // length 4, cap 5

	y := x[:2] // [a b]
	z := x[2:] // [c d]
	fmt.Println(x, len(x))
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(cap(x), cap(y), cap(z)) // 5 5 3

	y = append(y, "i", "j", "k")
	fmt.Println("x:", x) // [a b i j]
	fmt.Println("y:", y) // [a b i j k]
	fmt.Println("z:", z) // [i j]

	x = append(x, "x")
	fmt.Println("x:", x) // [a b i j x]
	fmt.Println("y:", y) // [a b i j x]
	fmt.Println("z:", z) // [i j]

	z = append(z, "y")
	fmt.Println("x:", x) // [a b i j y]
	fmt.Println("y:", y) // [a b i j y]
	fmt.Println("z:", z) // [i j y]

	a := make([]string, 0, 5)         // length - 0, cap - 5
	a = append(a, "a", "b", "c", "d") // [a b c d]
	b := a[:2:2]                      // origin[inicio:fim:limite (cap)]
	c := a[2:4:4]                     // a [2::4] - inválido

	// comprimento: fim - início
	// capacidade: limite - início

	fmt.Println(a, cap(a)) // [a b c d] 5
	fmt.Println(b, cap(b)) // [a b] 2
	fmt.Println(c, cap(c)) // [c d] 2

	// append cria outro array quando a capacidade não chega
}
