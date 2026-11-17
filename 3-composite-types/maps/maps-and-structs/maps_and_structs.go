package main

import "fmt"

func main() {
	intSet := map[int]struct{}{}

	vals := []int{1, 2, 3, 4, 5}

	for _, v := range vals {
		intSet[v] = struct{}{}
	}

	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}
}
