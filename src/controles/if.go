package main

import (
	"fmt"
	"math"
)

func sqrt(valor float64) string {
	if valor < 0 {
		return sqrt(-valor) + "i"
	}
	return fmt.Sprint(math.Sqrt(valor))
}

func main() {
	fmt.Println(sqrt(2), sqrt(4))
}
