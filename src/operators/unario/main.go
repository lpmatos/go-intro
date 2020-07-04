package main

import "fmt"

func main() {
	x, y := 1, 2

	// Apenas a pos-fix
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)
	y-- // y -= 1 ou y = y + 1
	fmt.Println(y)

}
