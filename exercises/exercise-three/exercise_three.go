package main

import (
	"fmt"
	"maps"
)

func main() {
	orcamento := 2000

	papelaria := map[string]int{
		"Caderno": 250,
		"Caneta":  120,
		"Pasta":   400,
	}

	compra := map[string]int{
		"Caderno": papelaria["Caderno"] * 3,
		"Caneta":  papelaria["Caneta"] * 2,
		"Pasta":   papelaria["Pasta"] * 1,
	}

	fmt.Println("Caderno:", papelaria["Caderno"])
	fmt.Println("Caneta:", papelaria["Caneta"])
	fmt.Println("Pasta:", papelaria["Pasta"])

	total := 0

	for i := range maps.Values(compra) {
		total += i
	}

	fmt.Println("Preço total:", total)
	fmt.Println("Dinheiro restante:", orcamento-total)
}
