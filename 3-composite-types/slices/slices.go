package main

import (
	"fmt"
	"slices" // A sequence of values.
)

func main() {
	var a = []int{10, 20, 30} // Using [...] makes an array. Using [] makes a slice.
	var b = []int{1, 5: 4, 6, 10: 100, 15}
	var d []int
	e := []int{1, 2, 3, 4}
	f := []int{1, 2, 3, 4}
	g := []string{"a", "b", "c"}
	h := []int{1, 2, 3}
	i := []int{}

	d = append(d, 1) // assign result to the variable that's passed in
	b = append(b, 16, 17, 18, 19, 20)

	i = append(i, h...)

	/*
		destino = append(slice de partida (recebe uma cópida das informações de i - localização dos dados, comprimento e capacidade)
		, elementos a acrescentar...)
	*/

	/*
		multidimensional slices

		var c [][]int
		c[0] = 10

		fmt.Println(c[2])
	*/

	fmt.Println(a)
	fmt.Println(b)                  // [1 0 0 0 0 4 6 0 0 0 100 15 16 17 18 19 20] - slice literal
	fmt.Println(d == nil)           // true
	fmt.Println(slices.Equal(e, f)) // true (same length and all of the elements are equal)
	fmt.Println(slices.Equal(e, a)) // false
	fmt.Println(len(d))             // 1
	fmt.Println(len(g))             // 3
	fmt.Println(i)                  // [1 2 3]

	// fmt.Println(slices.Equal(e, g)) does not compile - int != string

	/*
		A slice is the first type you've seen that ins't comparable.
		The only thing you can compare a slice with using == is nil.
	*/
}
