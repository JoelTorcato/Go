package main

import "fmt"

func main() {
	a := []string{"a", "b", "c", "d"}

	// destino = origin[inínio (começa no 0):fim(exluído, um a mais de onde quero parar)]
	b := a[:2]
	c := a[1:]
	d := a[1:3]
	f := a[:1]
	e := a[:]

	fmt.Println(a) // [a b c d]
	fmt.Println(b) // [a b]
	fmt.Println(c) // [b c d]
	fmt.Println(d) // [b c]
	fmt.Println(f) // a
	fmt.Println(e) // [a b c d]

	/*
		When you take a slice from a slice, you are not making a copy of the data.
		Instead, you now have two variables that are sharing memory. This means that changes
		to an element in a slice afftect all slices that share that element
	*/

	a[1] = "y"
	fmt.Println("a:", a) // a - [a y c d]
	fmt.Println("b:", b) // b - [a y]
	fmt.Println("c:", c) // c - [y c d]
	fmt.Println("f:", f) // f [a]

	b[0] = "x"
	fmt.Println("b:", b)  // b- [x y]
	fmt.Println("a:", a)  // a - [x y c d]
	fmt.Println("c: ", c) // c - [y c d]
	fmt.Println("f:", f)  // f [x]

	c[1] = "z"
	fmt.Println("c:", c) // c - [y z d]
	fmt.Println("a:", a) // a - [x z c d]
	fmt.Println("b:", b) // b - [x z]

}
