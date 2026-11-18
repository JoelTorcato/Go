package main

import "fmt"

func main() {
	papelaria := map[string]int{}

	papelaria["Caderno"] = 10
	papelaria["Caneta"] = 20
	papelaria["Borracha"] = 0

	papelaria["Caderno"] += 5
	papelaria["Caneta"] -= 3
	papelaria["Régua"] = 4

	v, ok := papelaria["Borracha"] // Atribuição

	fmt.Println(papelaria)      // map[Borracha:0 Caderno:15 Caneta:17 Régua:4]
	fmt.Println(len(papelaria)) // 4
	fmt.Println(v, ok)          // 0 true

	v, ok = papelaria["Mochila"] // Reatribuição
	fmt.Println(v, ok)           // 0 false
}
