package main

import "fmt"

func main() {
	lista := [4]float64{1, 2, 3, 4}
	fmt.Println(lista)
	var multiplicação float64 = lista[0] * lista[1] * lista[2] * lista[3]
	fmt.Println("O produto dos valores da lista é:", multiplicação)

}
