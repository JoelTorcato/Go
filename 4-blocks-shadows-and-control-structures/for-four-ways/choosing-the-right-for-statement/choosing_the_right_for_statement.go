package main

import "fmt"

func main() {
	// Most of the time, you're going to use the for-range format.

	/*
		When should you use the complete for loop?
		The best place for it is when you aren't iterating from the first element
		to the last element in a compound type.
	*/

	evenVals := []int{2, 4, 6, 8, 10}

	for i, v := range evenVals {
		if i == 0 {
			continue
		}
		if i == len(evenVals)-1 {
			break
		}
		fmt.Println(i, v) // 1 4  2 6  3 8
	}

	for i := 1; i < len(evenVals)-1; i++ {
		fmt.Println(i, evenVals[i]) // 1 4  2 6  3 8
	}

	/*
		If you want to skip over some of the runes in a string,
		you need to use a for-range loop so that it will properly process runes for you.
	*/

	// The body of the for loop should always contain a break or return because you'll rarely want to loop forever.
}
