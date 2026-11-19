package main

import "fmt"

func main() {
	var fila = []string{"Ana", "Bruno", "Carla"}

	atualizado := fila[1:]
	atualizado = append(atualizado, "Diogo")

	fmt.Println(fila)
	fmt.Println(atualizado)
	fmt.Println("Pessoas à espera:", len(atualizado))
	fmt.Println("Próximo a ser atendido:", atualizado[0])

}
