package main

import "fmt"

func main() {
	xArray := [4]int{1, 2, 3, 4}
	y := xArray[:2] // [1 2]
	z := xArray[2:] // [3 4]
	xArray[0] = 10
	xSlice := xArray[:]

	fmt.Println(xSlice) // [1 2 3 4]
	fmt.Println(xArray) // [10 2 3 4]
	fmt.Println(y)      // [10 2]
	fmt.Println(z)      // [3 4]

	/*
		Be aware that taking a slice from an array has the same memory-sharing properties as
		taking a slice from a slice.
	*/
}
