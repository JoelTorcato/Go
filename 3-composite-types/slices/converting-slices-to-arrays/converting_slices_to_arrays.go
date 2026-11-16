package main

import "fmt"

func main() {
	/*
		When you convert a slice to anarray, the data in the slice is copied to new memory.
		That means that changes to the slice won't affect the array, and vice versa.
	*/

	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	/*
		panic := [5]int(xSlice)
		fmt.Println(panic)
	*/

	xSlice[0] = 10

	fmt.Println(xSlice)     // [10 2 3 4]
	fmt.Println(xArray)     // [1 2 3 4]
	fmt.Println(smallArray) // [1 2]
}
