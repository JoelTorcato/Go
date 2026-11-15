package main

import "fmt"

func main() {
	var x int = 10
	var y byte = 100
	var sum3 int = x + int(y)
	var sum4 byte = byte(x) + y

	fmt.Println(sum3)
	fmt.Println(sum4)
}
