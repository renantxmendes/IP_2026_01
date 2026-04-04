package main

import  "fmt"

var nota float64
var conceito string

func main() {
	fmt.Print("Escreva a nota: ")
	fmt.Scan(&nota)

	if nota >= 9 && nota <= 10 {
		conceito = "A"
	} else
	if nota >= 7.5 && nota < 9 {
		conceito = "B"
	} else
	if nota >= 6 && nota < 7.5 {
		conceito = "C"
	} else {
		conceito = "D"
	}
		
	fmt.Printf("NOTA = %.2f CONCEITO = %s", nota, conceito)
}
