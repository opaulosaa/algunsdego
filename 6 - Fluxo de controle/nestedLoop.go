//Repetição hierárquica

package main

import "fmt"

func main() {
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora: ", horas)
		for minutos := 0; minutos < 60; minutos++ {
			fmt.Println("Minuto: ", minutos)
		}
	}

	//fazendo um "while":
	/*
		x := 0 -> declara uma variavel
		faz um for só com a condição -> for x < 10{
			fmt.Prinln("x é menor que dez")
			x++
		}
	*/
}
