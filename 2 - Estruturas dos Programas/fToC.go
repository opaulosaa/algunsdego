// ftoc é uma função de go que faz conversão de temperatura
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g ºF = %g º C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g º F = %g º C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) / 9
}

//a função específica pode vir depois da função main
