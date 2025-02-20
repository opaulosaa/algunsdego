package main

import "fmt"

var array [5]int

func main() {
	array[0] = 1
	array[1] = 9
	array[2] = 18
	fmt.Println(array[0], array[1], array[2])
	fmt.Println(array)

}
