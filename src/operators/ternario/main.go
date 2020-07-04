package main

import "fmt"

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}
func main() {
	// Não possui ternário
	fmt.Println(obterResultado(6.2))
}
