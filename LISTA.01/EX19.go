package main

import  "fmt"

var numeros int
var soma_num float64

func main() {
	fmt.Print("Digite um número positivo e maior que 1: ")
	fmt.Scan(&numeros)
	soma_num = 0
	for i := 1; i <= numeros ; i++ {
		soma_num = soma_num + (1/float64(i))
	}

	fmt.Printf("%.6f", soma_num)
}
