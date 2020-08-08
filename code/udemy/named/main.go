package main

import "fmt"

/*

Valores de retorno em Go podem ser nomeados e agirem como variáveis.

Esses nomes devem ser usados para documentar o significado dos valores de
retorno.

A declaração return sem argumentos retorna os valores atuais dos resultados.
Isso é conhecido como retorno limpo.

Instruções de retorno limpo devem ser usadas apenas em funções curtas.

*/

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	name := "Lucca"
	a, b := split(17)
	fmt.Println("Your name is:", name)

	fmt.Println("Resultado:", a, b)
}
