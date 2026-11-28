package main

import (
	"fmt"
	"strconv"
)

type opFuncType func(int, int) int

func sum(i, j int) int {
	return i + j
}

var opMap = map[string]opFuncType{
	"+": sum,
}

func main() {
	expressions := [][]string{ // slice de slice de strings
		{"10", "+", "5"},
	}

	/*
		expressions[0][1] "10" - O primeiro índice escolhe a expressão, e o segundo escolhe um elemento dentro dela.
	*/

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid statement: ", len(expression))
			continue
		}
		p1, err := strconv.Atoi(expression[0])

		if err != nil {
			fmt.Println(err)
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]

		/*
			Como op vale "+", Go procura a chave "+" em onMap.
			A entrada "+": sum associa essa chave ao valor da função sum.
			A consulta usa o comma-ok idiom e atribui dois valores.
		*/

		if !ok {
			fmt.Println("invalid operator: ", op)
			continue
		}

		p2, err := strconv.Atoi(expression[2])

		if err != nil {
			fmt.Println(err)
			continue
		}

		if p1 == p2 {
			fmt.Println("cannot have equal numbers:", p1)
			continue
		}

		result := opFunc(p1, p2)
		fmt.Println(result)
	}
}
