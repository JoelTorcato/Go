package main

import "fmt"

func main() {
	// A string is composed of a sequence of bytes
	var s string = "Hello there"
	var b byte = s[6]
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]

	var sol string = "Hello ☀️"
	var sol2 string = sol[4:7]
	var sol3 string = sol[:5]
	var sol4 string = sol[6:]

	fmt.Println(b)         // 116 (the UFT-8 value of a lowecase t)
	fmt.Println(s2)        // o t
	fmt.Println(s3)        // Hello
	fmt.Println(s4)        // there
	fmt.Println(sol)       // Hello ☀️
	fmt.Println(sol2)      // o � - That's because you copied only the first byte of the sun emoji's code point, which is not a valid code point on its own.
	fmt.Println(sol3)      // Hello
	fmt.Println(sol4)      // ☀️
	fmt.Println(len(sol))  // 12
	fmt.Println(len(sol4)) // 6

	/*
		Even though Go allows you to use slicing and indexing syntax with strings,
		you should use it only when you know that your string contains only characters that take up one byte
	*/
}
