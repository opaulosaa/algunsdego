package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //define duas variáveis (s e sep) do tipo string. Iniciam como duas strings vazias
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
