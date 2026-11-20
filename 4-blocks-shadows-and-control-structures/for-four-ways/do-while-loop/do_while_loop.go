package main

import "fmt"

func main() {
	/*
		for {
			things to do int the loop
			if !CONDITION {
				break
			}
		}
	*/

	a := 1
	for {
		fmt.Println("Hello") // One time
		if a != 0 {
			break
		}
	}
}
