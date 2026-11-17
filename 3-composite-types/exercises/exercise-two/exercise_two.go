package main

import "fmt"

func main() {
	var message string = "Hi 👩 and 👨"

	var rs []rune = []rune(message)

	fmt.Println(message)
	fmt.Println(string(rs[3]))
}
