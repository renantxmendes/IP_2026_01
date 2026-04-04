package main

import "fmt"

var n_1, n_2, n_3, res int

func main() {
	fmt.Print("Escreva 3 números inteiros com apenas uma unidade: ")
	fmt.Scan(&n_1, &n_2, &n_3)

	if n_1 == 0 || n_1 > 9 || n_2 > 9 || n_3 > 9 {
		fmt.Print("DÍGITO INVÁLIDO")
		return
	}

	res = n_1*100 + n_2*10 + n_3

	fmt.Printf("%d, %d", res, res*res)
}
