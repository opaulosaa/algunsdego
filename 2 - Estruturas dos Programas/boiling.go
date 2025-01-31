// exibe o ponto de ebulição da água
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) / 9
	fmt.Printf("water's boiling point = %g ºF or %g ºC\n", f, c)
}
