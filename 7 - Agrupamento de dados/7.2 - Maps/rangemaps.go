package main

import "fmt"

func main() {

	mapa := map[int]string{
		123: "muito legal",
		998: "menos legal",
		456: "massa",
		18:  "festa",
	}
	fmt.Println(mapa)

	for key, value := range mapa {
		fmt.Println(key, value)
	}

	fmt.Println("----------------")

	delete(mapa, 123)
	fmt.Println(mapa)

}
