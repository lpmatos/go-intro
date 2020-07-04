package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(100))

	// Inteiros sem sinal - uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println(b)
	fmt.Println("O byte é", reflect.TypeOf(b))

	// Inteiros com sinal - int8, int16, int32, int64
	i := math.MaxInt64
	fmt.Println(i)
	fmt.Println("O max inteiro de 64 é", reflect.TypeOf(i))

	// Rune - representação inteira da tabela unicode int32
	var c rune = 'c'
	fmt.Println(c)
	fmt.Println("O rune é", reflect.TypeOf(c))

	// Floats float32 e float64 (padrão)
	var salario float32 = 432.12
	fmt.Println("O salário é:", salario)
	fmt.Println("O tipo do salário é", reflect.TypeOf(salario))
	fmt.Println("O tipo do salário é", reflect.TypeOf(342.32))
	fmt.Println("O tipo do salário é", reflect.TypeOf(float32(342.32)))

	// Booleanos
	boolean := true
	fmt.Println("A variável boolean é:", boolean)
	fmt.Println("A variável boolean é:", !boolean)
	fmt.Println("O tipo da variável boolean é:", reflect.TypeOf(boolean))

	// String
	nome := "Lucca Pessoa da Silva Matos"
	fmt.Println("Seu nome é:", nome)
	fmt.Println("Tamanho:", len(nome))
	fmt.Println("Concatenando:", nome+"!")
	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))

	other := `
	Opa! Tudo bom??

	Prazer, me chamo Jarbes!
	`

	fmt.Println(other)

	// Go não tem char
	char := 'a'
	fmt.Println(char)
	fmt.Println(reflect.TypeOf(char))
}
