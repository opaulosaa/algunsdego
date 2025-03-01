/*
Crie uma slice usando make que possa conter todos os estados do Brasil.
Os estados: "Acre", "Alagoas", "Amapá", "Amazonas",
"Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso",

	"Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco",
	"Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia",
	 "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"

Demonstre o len e cap da slice.
Demonstre todos os valores da slice sem utilizar range.
*/
package main

import (
	"fmt"
)

func main() {
	estados := make([]string, 26, 26) //utiliza o make pra criar o slice, determinar cap e len

	//determinar os valores do slice
	estados = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo",
		"Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba",
		"Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte",
		"Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	//mostrar len e cap do slice
	fmt.Println("Comprimento:", len(estados))
	fmt.Println("Capacidade:", cap(estados))

	//mostrar todos os valores sem utilizar o range
	for i := 0; i < len(estados); i++ {
		fmt.Printf("Índice: %d, Estado: %s\n", i, estados[i])
	}
}
