// declarar pacote principal
package main

//importar função fmt
import "fmt"

//delcaração da variavel do ponto de ebulição da agua em F
const ebulicaoF = 212.0

// função principal
// := só pode ser usado dentro de um bloco de código
func main() {
	tempF := ebulicaoF

	tempC := (tempF - 32) * 5 / 9

	tempK := (tempC + 273.15)

	fmt.Printf("a temperatura de ebulição da agua em ºF = %g (%T), a temperatura de ebulição da agua em ºC = %g e a temperatura de ebulição da agua em K = %g ", tempF, tempF, tempC, tempK)
	//pode usar %g, %e ou %f para ponto flutuante, %d para inteiros e %s para strings
	//usa %T para descobrir o tipo, repetindo o nome a ser mostrado no final
}
