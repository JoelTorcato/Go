package main

import "fmt"

func main() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1

	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", smallI)
	fmt.Printf("%v\n\n", bigI)
}
