package main

import "fmt"

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}

	for i, v := range evenVals {
		fmt.Println(i, v)
		// i (index) -> Position in the data structure being iterated
		// v (value) -> Is the value at that position
	}
	/*
		If you don't need to access the key, use an underscore (_)
		as the variable's name. This tells Go to ignore the value.
	*/

	// Not print out the position

	for _, v := range evenVals {
		fmt.Println(v) // 2 4 6 8 10 12
	}

	/*
		What if your want the key but don't want the value?
		In this situation, Go allows you to just leave off the second variable
	*/

	uniqueNames := map[string]bool{
		"Fred":  true,
		"Raul":  true,
		"Wilma": true,
	}

	for k := range uniqueNames {
		fmt.Println(k) // Fred Raul Wilma
	} // The most common reason for iterating over the key is when a map is being used as a set.

	for _, v := range uniqueNames {
		fmt.Println(v) // true true true
	}
}
