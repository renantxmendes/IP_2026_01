package main

import f "fmt"

func main() {

	soma := 0.0
	num := 1.0
	baseDen := 15.0 
	sinal := 1.0

	f.Println("--- Calculando a Série passo a passo ---")

	for i := 1; i <= 15; i++ {
		den := baseDen * baseDen

		termoAtual := (num / den) * sinal
		soma = soma + termoAtual

		f.Printf("Termo %2d: %+.0f / %.0f \n", i, num*sinal, den)

		num = num * 2.0          
		baseDen = baseDen - 1.0  
		sinal = sinal * -1.0     
	}

	f.Printf("\nO valor total de S é: %f\n", soma)
}
