package main

import "fmt"

const PI = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

/*
Constantes são declaradas como variáveis, mas com a palavra chave const.

Constantes podem ser sequências de caracteres, booleanos ou valores numéricos.

Não podem ser declaradas utilizando a sintaxe :=

Constantes numéricas também são valores de alta precisão.

Quando declarada sem tipo, pega o tipo necessário pelo seu contexto.
*/

func needInt(x int) int           { return (x * 10) + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	const World = "世界"
	const Nome = "Lucca Pessoa da Silva Matos"
	fmt.Println("Hello", World)
	fmt.Println("Happy", PI, "Day")
	fmt.Println("Nome:", Nome)

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
