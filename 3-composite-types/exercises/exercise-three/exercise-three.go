package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	joel := Employee{
		"Joel",
		"Fernandes",
		1,
	}

	simao := Employee{
		firstName: "Simão",
		lastName:  "Fernandes",
		id:        2,
	}

	var susana Employee
	susana.firstName = "Susana"
	susana.lastName = "Torcato"
	susana.id = 3

	fmt.Println(joel)   // {Joel Fernandes 1}
	fmt.Println(simao)  // {Simao Fernandes 2}
	fmt.Println(susana) // {Susana Torcato 2}
}
