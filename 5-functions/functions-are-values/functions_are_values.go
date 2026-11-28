package main

import "fmt"

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func f3(a string) string {
	return a
}

func main() {
	var myFuncVariable func(string) int
	var returnString func(string) string
	myFuncVariable = f1

	result := myFuncVariable("Hello")
	fmt.Println(result) // 5

	myFuncVariable = f2
	result = myFuncVariable("Hello")
	fmt.Println(result)

	returnString = f3
	resultString := returnString("Hello")
	fmt.Println(resultString)

	/*
		The default zero value for a function variable is nil.
		Attempting to run a function variable with a nil value results in a panic.
	*/
}
