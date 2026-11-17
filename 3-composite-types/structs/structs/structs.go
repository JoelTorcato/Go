package main

import "fmt"

func main() {
	/*
		When you have related data that you want to group together,
		you should define a struct.
		A struct type that's defined within a function can be used only within that function
	*/

	type person struct { // A struct type is defined with the keyword type.
		name string
		age  int
		pet  string
	} // No commas separate the fields in a struct declaration.

	var fred person   // Since no value is assigned to fred, it gets the zero value for the person struct type.
	fmt.Println(fred) // { 0 }

	// A struct literal can be assigned to a variable as well
	bob := person{}
	fmt.Println(bob) // { 0 }

	julia := person{
		"Julia",
		18,
		"Dog",
	}

	/*
		This style has some advantages.
		It allows you to specify the fields in any order, and you don't need to provide a value for all fields.
		Any field not specified is set to its zero value (int - 0, string " ")
	*/
	joel := person{
		pet: "Cat",
		age: 17,
	}

	// You cannot mix the two struct literal styles.

	fmt.Println(julia) // {Julia 18 Dog}
	fmt.Println(joel)  // { 17 Cat}

	joel.name = "Joel"
	joel.pet = "Dog"

	fmt.Println(joel.name) // Joel (dotted notation for reading and writing to struct fields)
	fmt.Println(joel.pet)  // Dog
	fmt.Println(joel)      // {Joel 17 Dog}
}
