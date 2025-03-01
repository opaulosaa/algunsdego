/*
Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
"Nome", "Sobrenome", "Hobby favorito"
Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
package main

import "fmt"

func main() {

	//Criando um slice 2D de strings
	info := [][]string{
		{"paulosa", "leles", "tenis"},
		{"shimonzip", "messias", "jesus"},
		{"leozera", "preto", "livro"},
	}
	//imprimindo informações de cada linha e coluna
	for linhaIndex, linha := range info {
		for colunaIndex, valor := range linha {
			fmt.Printf("info[%d][%d] = %s\n", linhaIndex, colunaIndex, valor)
		}
	}
}
