package main

import (
	"fmt"
	"math/rand"
)

/*

- Cada programa Go é composto de pacotes.
- Programas sempre começam rodando pelo pacote **main**.
- Por convenção, o nome do pacote é o mesmo que o último elemento do caminho de importação.

*/

func main() {
	fmt.Println("My favorite number is:", rand.Intn(10))
}
