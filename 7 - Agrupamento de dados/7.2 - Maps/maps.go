package main

import "fmt"

func main() {
	amigos := map[string]int{ //criando map com chaves e valores (key, value)
		"shimonzip": 888999777,
		"leozera":   555444111,
	}
	fmt.Println(amigos)

	//adicionando novos keys e values ao map
	amigos["gopher"] = 111222333
	fmt.Println(amigos)

	//comma, ok -> saber se o valor existe ou nao no mapa, por exemplo:
	if existe, ok := amigos["fantasma"]; !ok {
		fmt.Println("não existe") //se existir valor, a saída será true
	} else {
		fmt.Println(existe)
	}

}
