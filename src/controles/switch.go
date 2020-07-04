package main

import (
	"fmt"
	"runtime"
)

/*
A instrução switch é uma forma mais curta de escrever uma sequência de declarações if - else.

Ela executa o primeiro caso cujo valor é igual à expressão da condição.
*/

func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}
