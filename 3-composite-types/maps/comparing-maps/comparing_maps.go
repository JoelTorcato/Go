package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]bool{
		"Xadrez":  true,
		"Futebol": false,
	}

	n := map[string]bool{
		"Xadrez":  true,
		"Futebol": false,
	}

	fmt.Println(m) // map[Futebol: false Xadrez: true]
	fmt.Println(n) // map[Futebol: false Xadrez: true]

	fmt.Println(maps.Equal(m, n)) // true
}
