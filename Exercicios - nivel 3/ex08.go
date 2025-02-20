// Crie um programa que utilize a declaração switch, sem switch statement especificado.
package main

import "fmt"

func main() {
	opcao := 2

	switch {
	case opcao == 1:
		fmt.Println("Abrir menu")
	case opcao == 2:
		fmt.Println("Escolher")
	case opcao == 3:
		fmt.Println("Confirmar")
	case opcao == 4:
		fmt.Println("Sair")
	}
}
