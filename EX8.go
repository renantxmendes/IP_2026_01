package main

import f "fmt"

var raio, altura, areaT, precoLata float64
var pi float64 = 3.14159


func main() {
	f.Print("Informe o raio e a altura da lata: ")
	f.Scan(&raio, &altura)
	areaT = 2*pi*raio*raio + 2*pi*raio*altura
	precoLata = areaT*100
	f.Printf("O VALOR DO CUSTO é de = %.2f\n", precoLata)
}
