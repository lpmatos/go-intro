package main

import (
	"fmt"
	"math"
)

/*

Em Go, um nome é exportado se ele começar com uma letra maiúscula.

Por exemplo, Pizza é um nome exportado, assim como Pi, que é exportado do pacote
math.

Ao importar um pacote, você pode referenciar apenas seus nomes exportados.

Todos os nomes não exportados não são acessíveis de fora do pacote.

*/

func main() {
	fmt.Println("PI value:", math.Pi)
}
