// array -> tem numero definido de elementos
// slice -> não tem número definido de elementos
package main

import "fmt"

func main() {

	novoArray := [5]int{1, 2, 3, 4, 5} //nome, operador curto, tamanho, tipo, elementos
	fmt.Println(novoArray)

	slice := []int{6, 7, 8, 9, 10}
	fmt.Println(slice)

	//inserir novos dados no slice
	slice2 := append(slice, 11)
	fmt.Println(slice2)

}
