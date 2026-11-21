package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_n!"}

outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				fmt.Printf("\n")
				continue outer
			}
		}
	}

palavra: // Notice that the label outer is indented by go fmt to the same level as the surrounding function.
	for _, sample := range samples {
		for i, r := range sample {
			if r == 'l' {
				break palavra
			}
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	/*
		Se eu usar continue eu passo para a próxima letra (inner loop).
		Se eu usar break eu termino o inner loop - executando ainda o Println() - e passo para o outer loop, que será utilizado para a segunda palavra.
		Se eu usar continue outer, eu pulo o resto das letras dessa palavra (se houver) e passo para a palavra seguinte, neste caso apple_n!
		Se eu usar break outer saiu dos dois ciclos e o programa lê o que vier a seguir.
	*/

	/*
		Nested for loops with labels are rare. They're most commonly used to implement
		algorithms similar to the following pseudocode.

			outer:
			for _, outerVal := range outerValues {
				for _, innerVal := range outerVal {
						// process innerVal
						if invalidSituation(innerVal) {
						continue outer
						}
				}
				here we have code that runs only when all of the innerVal values were successfully processed.
			}
	*/

}
