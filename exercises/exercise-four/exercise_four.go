package main

import "fmt"

func main() {
	type participante struct {
		primeiroNome  string
		ultimoNome    string
		idade         int
		identificador int
		confirmado    bool
	}

	utilizador := participante{
		primeiroNome:  "Álvaro",
		ultimoNome:    "Silva",
		idade:         20,
		identificador: 104,
		confirmado:    false,
	}

	utilizador.confirmado = true

	fmt.Println(utilizador.primeiroNome, utilizador.ultimoNome)
	fmt.Println(string(utilizador.primeiroNome[:2]))
	fmt.Printf("Idade: %d, Identificador: %d\n", utilizador.idade, utilizador.identificador)
	fmt.Println(utilizador.confirmado)
}
