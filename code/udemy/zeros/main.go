package main

import "fmt"

func main() {
	// Cada tipo tem um valor default - tipo zero

	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v\n%v\n%v\n%q\n%v", a, b, c, d, e)
}
