/*
Crie uma slice de tipo int
Atribua 10 valores a ela
Utilize range para demonstrar todos estes valores.
E utilize format printing para demonstrar seu tipo.
*/
package main

import "fmt"

func main() {
	//Cria slice com n valores
	meuSlice := []int{10, 20, 30, 40, 50, 1, 2, 3, 4, 5}

	//Itera o range para mostrar os valores e seus indices
	for i, v := range meuSlice {
		fmt.Printf("O indice %d está com o valor %d\n", i, v)
	}
	//format printing
	fmt.Printf("O tipo do slice é %T\n.", meuSlice)
}
