/*
Em package-level scope, atribua os seguintes valores às variáveis:
        1. para "x" atribua 42
        2. para "y" atribua "James Bond"
        3. para "z" atribua true
    2. Na função main:
        1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
        2. Demonstre a variável "s".
*/

package main

import "fmt"

var x2 = 42
var y2 = "James Bond"
var z2 = true

func main() {

	s := fmt.Sprintf("inteiro: %d, string: %s, bool: %t", x2, y2, z2) //Sprintf cria strings dinâmicos

	fmt.Println(s)
}
