package main

import "fmt"

func main() {
	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s), cap(s)) // [first, second, third] 3 3

	clear(s)
	fmt.Println(s, len(s), cap(s)) // [ ] 3 3

	// Remember, the zero value for a tring is an empty string ""!

	i := []int{1, 2, 3}
	fmt.Println(i, len(i), cap(i)) // [1 2 3] 3 3

	clear(i)
	fmt.Println(i, len(i), cap(i)) // [0 0 0] 3 3
}
