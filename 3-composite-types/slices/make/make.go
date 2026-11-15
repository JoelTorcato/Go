package main

import "fmt"

func main() {
	a := make([]int, 5) // creates and int slice with a length of 5 and a capacity of 5.
	// b := make([]int, 5, 10) length - 5, capacity - 10.
	c := make([]int, 0, 10)

	a = append(a, 10)
	c = append(c, 5, 6, 7, 8)

	fmt.Println(a)              // [0 0 0 0 0 10]
	fmt.Println(len(a), cap(a)) // 6 10
	fmt.Println(c)              // [5 6 7 8]
	fmt.Println(len(c), cap(c)) // 4, 10

	/*
		NEver specify a capacity that's less than the length! It is a compile time
		error to do so with a constant or numeric literal. If you use a variable
		to specify a capacity that's smallar than the length, your program will panic at runtime.
	*/
}
