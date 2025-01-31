// coleções de campos
// uteis para agrupar dados e formar registros
// type (algum nome, alguma variavel) struct
package main

import "fmt"

type pessoa struct {
	nome   string
	idade  int
	altura float64
}

func main() {
	fmt.Println(pessoa{"paulosa", 31, 1.74})

}
