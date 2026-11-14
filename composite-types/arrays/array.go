package main

import "fmt"

func main() {
	var a [3]int
	var b = [3]int{10, 20, 30}   // array literal
	var c = [...]int{10, 20, 30} // Inferir o comprimento da array apartir dos valores escritos
	var d = [12]int{1, 5: 4, 6, 7, 10: 100, 15}

	/*
		Go has only one-dimensional arrays

		var x = [2][3]int

		x[0] = 10
		fmt.Println(x[2])
		fmt.Println(len(x)) */

	var x = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	x[0][0] = 10
	fmt.Println(x[0])         // [10, 2, 3]
	fmt.Println(len(x), "\n") // 2 (array composto por outros arrays)

	fmt.Printf("%d\n", a)    // [0 0 0]
	fmt.Printf("%d\n", b[1]) // 20
	fmt.Println(b == c)      // Arrays are equal if they're the same lenght and contain equal values
	fmt.Println(b != c)      // false
	fmt.Printf("%d\n\n", d)  // [1, 0, 0, 0, 0, 4, 6, 7, 0, 0, 100, 15] (sparse array)
}
