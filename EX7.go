package main

import "fmt"

var temp_F, temp_C, pol, mil float64

func main() {
	fmt.Print("Digite a temperatura em Fharenheit: ")
	fmt.Scan(&temp_F)
	converterT()
	fmt.Print("Digite a quantidade de chuva em polegadas: ")
	fmt.Scan(&pol)
	mil = pol*25.4
	fmt.Printf("O VALOR EM CELSIUS = %.2f\nA QUANTIDADE DE CHUVA E = %.2f", temp_C, mil)
}

func converterT() {
	temp_C = 5*(temp_F - 32) / 9
}
