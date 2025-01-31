package main

import "fmt"

func main() {
	/*arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	slice := arr[2:5]
	fmt.Println()*/

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

}

//var x []float 64 -> declara a variavel como um array e fala o tipo que estará no array
// Slice(fatia) -> slice := make([]float64, 4) obs: o 4, no caso, é o tamanho da fatia do array
//slice := [low:high]
//slice := arr[1:4]
//Append: acrescentar elementos
//copiar uma fatia para outra -> copy
