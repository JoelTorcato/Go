package main

import "fmt"

func main() {
	totalWins := map[string]int{}

	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2

	fmt.Println(totalWins["Orcas"])   // 1
	fmt.Println(totalWins["Kittens"]) // 0 - When you try to read the value assigned to a map key that was never set, the map returns the zero value for the map's value type

	totalWins["Kittens"]++

	fmt.Println(totalWins["Kittens"]) // 1

	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"]) // 3

	people := map[int]string{}

	people[1] = "Joel"
	people[2] = "Simão"

	fmt.Println(people)
	fmt.Println(people[1]) // Joel
	fmt.Println(people[2]) // Simão
	fmt.Println(people[3]) //

	people[3] = "Mãe"

	fmt.Println(people[3]) // Mãe
}
