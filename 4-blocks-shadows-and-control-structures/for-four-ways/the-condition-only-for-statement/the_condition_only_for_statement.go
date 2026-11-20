package main

import "fmt"

func main() {
	// Like while loop

	i := 1

	for i < 100 {
		fmt.Println(i) // 1 2 4 8 16 32 64
		i *= 2
	}
}
