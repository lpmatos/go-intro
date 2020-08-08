package main

import "fmt"

// A declaração defer adia a execução de uma função até o final do retorno da função.

// Os argumentos das chamadas adiadas são avaliados imediatamente, mas a função chamada não é executada até o retorno da função.

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
