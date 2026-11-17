package main

import "fmt"

func main() {
	/*
	 A set is a data type that ensures there is at most one of a value,
	 but doesn't guarantee that the values are in any particular order.
	*/

	intSet := map[int]bool{} // Create a map where the keys are of int type and the values are of bool type

	vals := []int{5, 10, 2, 8, 7, 3, 100} // We cannot have duplicate keys in a map.
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(vals[0])     // 5
	fmt.Println(len(vals))   // 7
	fmt.Println(len(intSet)) // 7
	fmt.Println(intSet[5])   // true - If number 5 exists
	fmt.Println(intSet[200]) // false

	if intSet[100] {
		fmt.Println("100 is in the set")
	}

}
