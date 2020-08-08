package main

import "fmt"

/*
Go tem apenas uma estrutura de laço, o for.

Possui 3 componentes separados por ponto e vírgula:

* A instrução inicial: executada antes da primeira iteração.
* A expressão de condição, avaliada antes de cada iteração.
* A pós instrução: executada no final de cada iteração.

A instrução inicial será frequentemente uma declaração de uma variável curta.

Essas variáveis curtas são visíveis apenas no escopo do for.
*/

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	accumlador := 1
	const TAM = 100
	for i := 1; i <= (TAM + 1); i++ {
		if (i % 2) == 0 {
			accumlador++
		} else {
			fmt.Println("Skip:", i)
			continue
		}
	}
	fmt.Println(accumlador)
}
