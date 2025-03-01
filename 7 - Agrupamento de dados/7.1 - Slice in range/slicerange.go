package main

import "fmt"

func main() {
	slice := []string{"banana", "maça", "jaca"}

	for index, valor := range slice {
		fmt.Println("no indice ", index, "temos o valor ", valor)
	}
}
