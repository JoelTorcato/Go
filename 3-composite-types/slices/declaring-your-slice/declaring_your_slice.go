package main

import "fmt"

func main() {
	// var data []int - Declaring a slice that might stay nil
	// var x = []int{} - You can create a slice using an empty slice literal

	data := []int{2, 4, 6, 8} // Declaring a slice with default values

	fmt.Println(data)

	/*
		If you have a good idea of how large your slice needs to be,
		but don't know what those values will be when you're writting
		the program, use make.
	*/
}
