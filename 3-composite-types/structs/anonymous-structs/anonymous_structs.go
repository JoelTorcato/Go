package main

import "fmt"

func main() {
	var person struct { // anonymous struct
		name string
		age  int
		pet  string
	}

	person.name = "Joel"
	person.age = 17
	person.pet = "Cat"

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(person)      // {Joel 17 Cat}
	fmt.Println(pet)         // {Fido dog}
	fmt.Println(pet.name)    // Fido
	fmt.Println(person.name) // Joel
}
