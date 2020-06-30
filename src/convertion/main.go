package main

import (
	"fmt"
	"math"
)

/*
A expressão T(v) converte o valor v para o tipo T.

Inferência de tipo.

Ao declarar uma variável, sem especificar o seu tipo (usando var sem um tipo
ou na sintaxe :=), o tipo da variável é inferida a partir do valor ao lado
direito.

Quando o lado direito da declaração é digitado, a nova variável é do mesmo tipo.

Porém, quando o lado direito contém uma constante numérica não tipificada, a
nova variável pode ser int, float64 ou complex128 dependendo da precisão da
constante.

*/

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	var i int = 10
	j := i
	fmt.Println(i, j)
	v := 42 // change me!
	fmt.Printf("v is of type %T - value %v\n", v, v)
}
