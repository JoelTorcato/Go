package main

import (
	"fmt"
	"math/rand"
)

func main() {
	slice := make([]int, 0)

	for i := 0; i <= 10; i++ {
		slice = append(slice, rand.Intn(10))
	}
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		switch {
		case slice[i]%2 == 0 && slice[i]%3 == 0:
			fmt.Println("Six")
		case slice[i]%2 == 0:
			fmt.Println("Two")
		case slice[i]%3 == 0:
			fmt.Println("Three")
		default:
			fmt.Println("Never mind")
		}
	}
}
