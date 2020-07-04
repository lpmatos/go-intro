package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.\n\n")

	fmt.Println("Linha")
	fmt.Println("Diferente")

	x := 32.321
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é", x)
	fmt.Println("O valor de x é " + xs)
	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 12.212121
	c := false
	d := "Lucca"

	fmt.Printf("\n%d %f %.1f %v %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
