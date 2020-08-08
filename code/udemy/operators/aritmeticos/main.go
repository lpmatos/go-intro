package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2
	c := 3.0
	d := 2.0

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Divisão =>", float64(a)/float64(b))
	fmt.Println("Módulo =>", a%b)

	// Bitwise - Pera o número, converte para binário e depois opera bit a bit segundo o operador.
	fmt.Println("E =>", a&b)   // 11&10 = 10
	fmt.Println("Ou =>", a|b)  // 11|10 = 11
	fmt.Println("Xor =>", a^b) // 11^10 = 01

	// Math
	fmt.Println("Maior =>", math.Max(c, d))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
