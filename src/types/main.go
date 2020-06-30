package main

import (
	"fmt"
	"math/cmplx"
)

/*
Os tipos básicos em Go são:

* bool
* string
* int, int8, int32, int64
* uint, uint8, uint32, uint64, uintptr
* byte - uint8
* rune - int32
* float32 float64
* complex64 complex 128

Os tipos int, uint e uintptr são geralmente de 32 bits em sistemas de 32 bits e 64 bits em sistemas de 64 bits.

Quando você precisar de um valor inteiro deverá usar int, a menos que tenha um motivo específico para usar um tipo de inteiro com tamanho especificado ou sem sinal.

Variáveis declaradas sem um valor inicial explicitado darão seu valor zero.

O valor zero é:

* 0 para tipos numéricos.
* false para tipos booleanos
* "" (string vazia) para strings

*/

var (
	MaxInt uint64     = 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	name := "Lucca Pessoa"
	ToBe := true
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", name, name)
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
