package main

import (
	 "fmt"
	  "math"
)


var aresta, area_B, alt, vol float64

func main() {
	fmt.Print("Digite o valor da altura e da aresta da base da pirâmide: ")
	fmt.Scan(&alt, &aresta)

	area_B = ((3*aresta*aresta)*math.Sqrt(3))/2
	vol = area_B*alt/3

	fmt.Printf("O VOLUME DA PIRÂMIDE E = %.2f METROS CÚBICOS", vol)
}
