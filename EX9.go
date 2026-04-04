package main

import f "fmt"

var a, b, c, delta float64

func main() {
	f.Print("Escreva os coeficientes a, b e c: ")
	f.Scan(&a, &b, &c)

	delta = (b*b) - (4*a*c)

	f.Printf("O RESULTADO DE DELTA E = %.2f", delta)
}
