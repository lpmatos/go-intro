package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Igualdade == ->", "banana" == "banana")
	fmt.Println("Diferença != ->", "banana" != "banana")
	fmt.Println("Maior > ->", 3 > 2)
	fmt.Println("Menor > ->", 3 < 2)
	fmt.Println("Maior ou Igual >= ->", 3 >= 2)
	fmt.Println("Menor ou Igual <= ->", 3 <= 2)
	fmt.Println("Negação ->", !(1 == 1))

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println(d1, d2)
	fmt.Println("Datas ->", d1 == d2)
	fmt.Println("Datas ->", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Lucca Pessoa"}
	p2 := Pessoa{"Lucca Pessoa"}

	fmt.Println(p1.Nome == p2.Nome)
	fmt.Println(p1.Nome != p2.Nome)
}
