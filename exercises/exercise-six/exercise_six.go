package main

import "fmt"

func main() {
	agenda := map[string]int{
		"Ana":   912345678,
		"Bruno": 923456789,
		"Carla": 934567890,
	}

	agenda["Bruno"] = 965432100
	delete(agenda, "Carla")
	agenda["Diogo"] = 945678901

	numero, confirmacao := agenda["Ana"]

	fmt.Println(agenda, len(agenda)) // map[Ana:912345678 Bruno:965432100 Diogo:945678901]
	fmt.Println(numero, confirmacao) // 912345678, true

	numero, confirmacao = agenda["Carla"]
	fmt.Println(numero, confirmacao) // 0, false
}
