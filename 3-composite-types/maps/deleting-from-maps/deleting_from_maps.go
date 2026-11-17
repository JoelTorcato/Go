package main

import "fmt"

func main() {
	m := map[string]int{
		"Joel":  0,
		"Simão": 1,
	}

	fmt.Println(m)      // map[Joel:0 Simão:1]
	fmt.Println(len(m)) // 2

	delete(m, "Joel")   // delete(map, key), dele the key-value
	fmt.Println(m)      // map[Joel:0]
	fmt.Println(len(m)) // 1

	/*
		If the key isn't present in the map or if the map is nil,
		nothing happens. The delete function doesn't return a value
	*/
}
