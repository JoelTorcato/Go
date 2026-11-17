package main

import "fmt"

func main() {
	m := map[string]int{
		"Joel":  0,
		"Simão": 1,
	}

	fmt.Println(m)      // map[Joel:0 Simão:1]
	fmt.Println(len(m)) // 2

	delete(m, "Joel")
	fmt.Println(m)      // map[Joel:0]
	fmt.Println(len(m)) // 1
}
