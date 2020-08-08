package main

import "fmt"

// O pacote fmt implementa E/S formatada com funções análogas ao printf e scanf do C. O formato "verbos" são derivados de C, mas são mais simples.

// Para mais informações - https://golang.org/pkg/fmt/

func main() {
	fmt.Println("Hello Linux Academy!")

	name := "Lucca Pessoa da Silva Matos"
	age := 20

	fmt.Printf("Go representation value - name: %#v\n", name)
	fmt.Printf("Go representation value - age: %#v\n", age)

	fmt.Printf("Go representation type of value - name: %T\n", name)
	fmt.Printf("Go representation type of value - age: %T\n", age)
}
