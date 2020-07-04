package main

import "fmt"

// Struct é uma coleção de campos.

// Campos (atributos) struct são acessados através do ponto

// Ponteiros para structs - campos podem de structs podem ser acessados através de um ponteiro de estrutura.

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v)

	// Criando ponteiro para o endereço de memória da struct v
	p := &v

	// Acessando o campo X dessa struct (*p).X
	p.X = 1e9
	fmt.Println(v)
}
