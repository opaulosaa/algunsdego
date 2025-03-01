/*
Comece com essa slice:
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
Utilizando slicing e append, crie uma slice y que contenha os valores:
[42, 43, 44, 48, 49, 50, 51]
*/
package main

import "fmt"

func main() {
	// 0   1    2   3   4   5   6   7   8   9
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51} //cria um slice de inteiros
	fmt.Println(x)

	//Utilizando slicing e append, crie uma slice y
	//x = append(x[:3], x[len(x) - 4:]... ) -> pode ser assim também, utilizando a função len(x)
	x = append(x[:3], x[:6]...)
	fmt.Println(x)
}
