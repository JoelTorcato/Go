package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_n!"}

	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r)) // i -> index, r - the numeric value of the letter, string(r) -> numeric value of the letter type converted to a string
		}
		fmt.Println()
	}
}
