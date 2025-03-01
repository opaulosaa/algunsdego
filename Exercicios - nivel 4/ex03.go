/*
Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
Do quinto ao último item do slice (incluindo o último item!)
Do segundo ao sétimo item do slice (incluindo o sétimo item!)
Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
*/
package main

import "fmt"

func main() {
	//Cria slice com n valores
	meuSlice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(meuSlice[:3])
	fmt.Println(meuSlice[4:])
	fmt.Println(meuSlice[1:7])
	fmt.Println(meuSlice[2:9])
	fmt.Println(meuSlice[2 : len(meuSlice)-1])
}
