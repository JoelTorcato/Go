package main

import "fmt"

func main() {
	var a rune = 'x' // guarda em valor binário
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)

	fmt.Println(s)        // x
	fmt.Println(s2)       // y
	fmt.Println(a)        // decimal - 120
	fmt.Printf("%x\n", a) // hexadecimal - 78
	fmt.Printf("%b\n", a) // binário - 01111000

	/*
		As of Go 1.15, go vet blocks a type conversion to string from any integer type other
		than rune or byte
	*/

	var v string = "Hello, ☀️"
	var bs []byte = []byte(v)
	var rs []rune = []rune(v) // Slices of runes are uncommon

	fmt.Println(v)  // Hello, ☀️
	fmt.Println(bs) // string converted to UTF-8 bytes
	fmt.Println(rs) // string converted to runes
}
