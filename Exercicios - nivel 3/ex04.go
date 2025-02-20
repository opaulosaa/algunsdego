//Crie um loop utilizando a sintaxe: for {}
//Utilize-o para demonstrar os anos desde que você nasceu.

package main

import "fmt"

func main() {
	ano := 1993
	anoAtual := 2025

	for {
		if ano > anoAtual {
			break
		}
		fmt.Println(ano)
		ano++
	}
}
