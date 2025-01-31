package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 8
	x[1] = 5.4
	x[2] = 7.45
	x[3] = 4.9
	x[4] = 9

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i] // soma o total com os elementos do array
	}

	fmt.Println(total / float64(len(x)))
}
