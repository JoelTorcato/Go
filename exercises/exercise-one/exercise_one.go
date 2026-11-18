package main

import "fmt"

func main() {
	material := []string{"Caderno", "Caneta", "Borracha"}
	maria := make([]string, len(material))
	copy(maria, material[:])

	maria = append(maria, "Régua")
	maria[0] = "Bloco de notas"

	fmt.Println(material, len(material))
	fmt.Println(maria, len(maria))
}
