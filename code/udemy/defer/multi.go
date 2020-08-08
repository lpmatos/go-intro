package main

import "fmt"

// Empilhando defers... Chamdas de funções adiadas são empuradas por uma pilha.

// Quando a função retorna, as chamadas de adianto são executadas na ordem em que a última a entrar é a primeira a sair.

// Uma declaração defer(de adiamento) envia uma chama de função para uma lista. A lista de chamadas salvas é executadas após o retorno da função circundante.

// O adianmento é comumento usado para simplificar funções que executam várias ações de limpeza.

// Os argumentos de uma função adiada são avaliados quando a instrução adiada é avaliada.

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
