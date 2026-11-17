package main

import "fmt"

func main() {
	// A cleared map has its length set to zero, unlike a cleared slice.

	m := map[int]string{
		1: "Joel",
		2: "Simão",
	}

	fmt.Println(m)      // map[1:Joel 2:Simão]
	fmt.Println(len(m)) // 2

	clear(m)            // m = map name
	fmt.Println(m)      // map[]
	fmt.Println(len(m)) // 0
}
