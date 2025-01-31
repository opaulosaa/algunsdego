// função associada a um tipo particular
// em POO, objeto é um valor(variavel) e método
// é uma função associada a esse objeto
package main

import "fmt"

type retangulo struct {
	base, altura float64
}

func (r *retangulo) area() float64 {
	return r.base * r.altura
}

func (r retangulo) perimetro() float64 {
	return 2*r.base + 2*r.altura
}

func main() {
	r := retangulo{base: 10, altura: 5}

	fmt.Println("Área:", r.area())
	fmt.Println("Perimetro:", r.perimetro())
}
