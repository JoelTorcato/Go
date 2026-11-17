package main

import "fmt"

func main() {
	greeting := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	f := greeting[:2]
	s := greeting[1:4]
	t := greeting[3:]

	fmt.Println(greeting) // [Hello Hola नमस्कार こんにちは Привіт]
	fmt.Println(f)        //[Hello Hola]
	fmt.Println(s)        // [Hola नमस्कार こんにちは]
	fmt.Println(t)        // [こんにちは Привіт]
}
