package main

import "fmt"

func main() {
	xSlice := []int{1, 2, 3, 4}
	xArrayPointer := (*[4]int)(xSlice)

	/*
		After converting a slice to an array pointer, the storage between the two is shared.
		A change to one will change the other.
	*/

	xSlice[0] = 10
	xArrayPointer[1] = 20

	fmt.Println(xSlice)        // [10 20 3 4]
	fmt.Println(xArrayPointer) // &[10 20 3 4]
}
