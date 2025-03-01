/*
Crie um map com key tipo string e value tipo []string.
Key deve conter nomes no formato sobrenome_nome
Value deve conter os hobbies favoritos da pessoa
Demonstre todos esses valores e seus índices.
*/
package main

import "fmt"

func main() {

	info := map[string]string{
		"alencar_paulo": "tenis",
		"shimon_natan":  "jesus",
	}
	fmt.Println(info)
	fmt.Println("----------")

	for key, value := range info {
		fmt.Println(key, value)
	}

}
