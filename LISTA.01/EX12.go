package main

import  "fmt"

var qnt_horas, preco int

func main() {
	fmt.Print("Digite a quantidade de horas: ")
	fmt.Scan(&qnt_horas)
	
	preco = ((qnt_horas - qnt_horas%3)/3)* 10 + qnt_horas%3*5
	
	fmt.Printf("O VALOR A PAGAR E = %.2d", preco)
}
