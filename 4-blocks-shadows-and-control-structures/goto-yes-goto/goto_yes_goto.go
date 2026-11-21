package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		In Go, a goto statement specifies a labeled line of cade, and execution jumps to it.
		However, you can't jump anywhere. Go forbids jumps that skip over variable
		declarations and jumps that go into an inner or parallel block.

		a := 10
			goto skip
			b := 10

		skip:
			c := 30
			fmt.Println(a, b, c)

			if c > a {
				goto inner
			}
			if a < b {
			inner:
				fmt.Println("a is less than b")
			}
	*/

	a := rand.Intn(10)

	for a < 100 {
		if a%5 == 0 {
			goto done // serve para saltar diretamente para um ponto identificado por um rótulo (label) na mesma função
		}
		a = a*2 + 1
	}
	fmt.Println("do something when the loop completes normally")

done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

	/*
		You should try very hard to avoid using goto. But in the rare
		situations where it makes your code more readable, it is an option.
	*/
}
