package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Cuidado " + string(123))

	// Int para string
	fmt.Println("Valida " + strconv.Itoa(123))
	fmt.Println("Valida " + fmt.Sprint(123))

	// String to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num)
	fmt.Println(reflect.TypeOf(num))

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println(b)
		fmt.Println(reflect.TypeOf(b))
	}
}
