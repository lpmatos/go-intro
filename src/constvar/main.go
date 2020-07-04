package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.14124
	var raio = 3.2
	area := PI * m.Pow(raio, 2)
	fmt.Println("A = ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, true

	fmt.Println(e, f)

	g, h, i := 2, "Lucca", false

	fmt.Println(g, h, i)
}
