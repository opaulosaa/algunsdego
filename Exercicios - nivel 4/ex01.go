/*
Crie um array que suporte 5 valores to tipo int
Atribua valores aos seus índices
Utilize range e demonstre os valores do array.
Utilizando format printing, demonstre o tipo do array.
*/
package main

import "fmt"

func main() {
	//Criação de array com 5 inteiros
	meuArray := [5]int{10, 20, 30, 40, 50}

	//Demonstra os valores do array em cada posicao
	for index, valor := range meuArray {
		fmt.Printf("Índice: %d, Valor: %d\n", index, valor)
	}
	//utiliza o format printing
	fmt.Printf("O tipo do array é: %T\n", meuArray)

}
