package main

import "fmt"

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
