//Crie um loop utilizando a sintaxe: for condition {}
//Utilize-o para demonstrar os anos desde que você nasceu.

package main

import "fmt"

func main() {
	ano := 1993

	for ano >= 1993 {
		fmt.Println(ano)
		ano++
		if ano >= 2025 {
			break
		}
	}
}
