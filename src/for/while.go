package main

import "fmt"

func main() {
	sum := 1
	fmt.Println("Exemplo de uso do for como while em Go...")
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Resultado ->", sum)
}
