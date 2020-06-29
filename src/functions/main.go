package main

import "fmt"

/*

Uma função pode ter zero ou mais argumentos.

Todo argumento deve possuir uma tipagem.

Observe que o tipo vem após o nome da variável.

Quando dois ou mais parâmetros nomeados consecutivos de uma função compartilham
o mesmo tipo, você pode omitir o tipo de todos menos o último.

Uma função pode retornar resultados múltiplos. Isso lembra um pouco Python, que
ao retornar vários valores separados por vírgula, na verdade você recebe uma
tupla com esses valores.
*/

func soma(x int, y int) int {
	fmt.Println("Somando valores...")
	return x + y
}

func swap(x, y string) (string, string, string) {
	fmt.Println("Swap strings...")
	return y, x, "default"
}

func add(x, y int) int {
	fmt.Println("Adicionando valores...")
	return x + y
}

func main() {
	fmt.Println("Resultado:", soma(42, 13))
	fmt.Println("Resultado:", add(4325, 13))
	a, b, c := swap("hello", "world")
	fmt.Println("Strings:", a, b, c)
}
