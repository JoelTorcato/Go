package main

import "fmt"

func main() {
	/*
		The comma ok idiom is used in Go when you want to differentiate
		between reading a value and getting back the zero value.

	*/
	m := map[string]int{
		"hello": 5,
		"world": 1,
	}
	fmt.Println(m["hello"]) // 5

	v, ok := m["hello"]
	/*
		The first gets the value (v) associated with the key.
		The second value returned is a bool (ok).
	*/
	fmt.Println(v, ok) //  5 true

	v, ok = m["world"]
	fmt.Println(v, ok) // 1 true

	v, ok = m["goodbye"]
	fmt.Println(v, ok) // 0 false
}
