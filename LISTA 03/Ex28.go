package main

import f "fmt"
import m "math"

func main() {
	soma := 0.0
	sinal := 1.0
	denominadorBase := 1.0

	for i := 1; i <= 51; i++ {
		termoAtual := (1.0 / m.Pow(denominadorBase, 3.0)) * sinal
		
		soma = soma + termoAtual
		sinal = sinal * -1.0
		denominadorBase = denominadorBase + 2.0
	}

	piAproximado := m.Cbrt(soma * 32.0)

	f.Printf("Valor aproximado de PI com 51 termos: %.15f\n", piAproximado)
}
