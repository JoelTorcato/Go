package main

import "fmt"

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func mutiplicateTo(base float64, vals ...float64) []float64 {
	out := make([]float64, 0, len(vals))

	for _, v := range vals {
		out = append(out, base*v)
	}
	return out
}

func main() {
	fmt.Println(addTo(3))                     // []
	fmt.Println(addTo(3, 2))                  // [5]
	fmt.Println(addTo(3, 2, 4, 6, 8))         // [5 7 9 11]
	fmt.Println(mutiplicateTo(1, 2, 3, 4, 5)) // [2 3 4 5]

	a := []int{4, 3}
	fmt.Println(addTo(3, a...))                    // [7 6]
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // [4 5 6 7 8]
	// You must put three dots (...) after the variable or slice literal. If you do not, it is a compile-time error.

	b := []float64{4, 3, 2}
	fmt.Println(mutiplicateTo(5, b...))                     // [20 15 10]
	fmt.Println(mutiplicateTo(2, []float64{10, 20, 30}...)) // [20 40 60]
}
