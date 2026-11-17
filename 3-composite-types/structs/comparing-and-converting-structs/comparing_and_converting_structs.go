package main

import "fmt"

func main() {
	type firstPerson struct {
		name string
		age  int
	}

	f := firstPerson{
		name: "Joel",
		age:  17,
	}

	var g struct {
		name string
		age  int
	}

	g = f
	fmt.Println(f == g) // true
}
