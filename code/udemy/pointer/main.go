package main

import "fmt"

// Go tem ponteiros. Um ponteiro guarda na memória o endereço de uma variável.
// O operador & gera um ponteiro para seu operando.
// O operador * indica valor subjacente do ponteiro.
// Isto é conhecido como "desreferenciamento" ou "indirecionamento".

func main() {
	i := 1
	fmt.Println(i, &i)

	// Go não possui aritmética de ponteiros
	fmt.Println("Criando ponteiro")
	var p *int = nil

	fmt.Println(p, &p)
	fmt.Println("Associando endereço de memória da variável i ao ponteiro p")
	p = &i
	fmt.Println(p, &p, *p)

	*p++

	fmt.Println(i, *p)
}
