package main

import "fmt"

func main() {
	x := 4

	if x == 2 || x == 3 {
		fmt.Println("x = 2 ou x = 3")
	} else if x == 4 {
		fmt.Println("x = 4")
	} else {
		fmt.Println("erroooooo")
	}

}
