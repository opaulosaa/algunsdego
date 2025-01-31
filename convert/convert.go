package main

import "fmt"

var a int = 20
var b string = "nome"

func main() {
	var c float64 = float64(a)
	fmt.Printf("o valor de c é: %g  e o valor de b é: %s ", c, b)
}
