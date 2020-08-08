package main

import "fmt"

/*

A instrução var declara uma lista de variáveis, o tipo é o último a ser passado.

Essa declaração pode estar num pacote ou a nível de função.

A declaração var pode incluir inicializadores, uma por variável.

Se um inicializador está presente, o tipo pode ser omitido... A variável terá o
tipo do inicializador.

Dentro de uma função, ao invés de utilizar o var, você pode utilizar a
instrução de atribuição curta := com o tipo implícito.

Fora de uma função cada estrutura comça com uma palavra chave (var, func...) e
não é possível usar o :=

*/

var c, python, java = true, false, "nooo!"

func main() {
	var i int = 10
	idade := 100
	fmt.Println(i, c, python, java, idade)
}
