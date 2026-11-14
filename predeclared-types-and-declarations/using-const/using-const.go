package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "Hello"

	fmt.Println(x)
	fmt.Println(y, "\n")

	// x = x + 1 - this will not compile!
	// y = "Bye" - this will not compile!
}
