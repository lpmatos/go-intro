package main

import "fmt"

func aprovado(nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado com nota ->", nota)
	} else {
		fmt.Println("Reprovado com nota ->", nota)
	}
}

func nota(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 5 && nota < 8 {
		return "C"
	} else if nota >= 3 && nota < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	aprovado(7.5)
	aprovado(5.1)
	fmt.Println("Seu conceito é -> ", nota(4.0))
	fmt.Println("Seu conceito é -> ", nota(6.0))
	fmt.Println("Seu conceito é -> ", nota(8.3))
	fmt.Println("Seu conceito é -> ", nota(9.3))
	fmt.Println("Seu conceito é -> ", nota(2.3))
}
