package main

import "fmt"

func main() {
	// map[keyType]valueType

	/* You can use a var declaration to create a map variable that's set to its zero value:
	var nilMap map[string]int // (string keys and int values)

	You can use a := declaration to create a map variable by assigning it a map literal
	totalWins := map[string]int{} // != nil map
	*/

	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}

	// fmt.Println(nilMap == nil) // true
	fmt.Println(teams)
	fmt.Println(len(teams))   // 3
	fmt.Println(teams == nil) // false

	/*
		If you know how many key-value pairs you intend to put in the map but don't know
		the exact values, you can use make to create a map with a default size
		ages := make(map[int][]string, 10) // length 0
	*/
}
