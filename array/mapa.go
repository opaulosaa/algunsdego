//coleção ordenada (lista) formada por pares chaves-valor
//var x map[string]int
// x é um mapa de strings para ints

package main

import "fmt"

func main() {

	/*x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])*/

	/*x := make(map[int]int)
	x[1] = 20
	fmt.Println(x[1])*/

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["O"] = "Oxigênio"
	elemento["N"] = "Nitrogênio"
	fmt.Println(elemento["H"], elemento["O"], elemento["N"])

}
