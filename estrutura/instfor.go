package main

import "fmt"

func main() {
	/*for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Número Par")
		} else {
			fmt.Println("Número Ímpar")
		}
		fmt.Println(i)
	}*/

	for i := 0; i < 5; i++ {
		switch i {
		case 0:
			fmt.Println("zero")
		case 1:
			fmt.Println("um")
		case 2:
			fmt.Println("dois")
		case 3:
			fmt.Println("tres")
		case 4:
			fmt.Println("quatro")

		}
	}
}
