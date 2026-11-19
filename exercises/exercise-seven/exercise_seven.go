package main

import (
	"fmt"
	"slices"
)

func main() {
	var aluno = []float64{12, 15, 10}
	var simulado = make([]float64, len(aluno))

	copy(simulado, aluno[:])
	simulado[2] = 16

	media := (aluno[0] + aluno[1] + aluno[2]) / 3
	mediaSimulado := (simulado[0] + simulado[1] + simulado[2]) / 3

	fmt.Println(aluno)                         // [12 15 10]
	fmt.Println(simulado)                      // [12 15 16]
	fmt.Println(media)                         // 12.333333333333334
	fmt.Println(mediaSimulado)                 // 14.333333333333334
	fmt.Println(slices.Equal(aluno, simulado)) // false

}
