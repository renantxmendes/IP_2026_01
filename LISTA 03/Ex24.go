package main

import f "fmt"
import m "math"
func main() {
	var a float64
	f.Print("Digite o valor do ângulo A (em radianos): ")
	f.Scan(&a)

	senA := 0.0
	sinal := 1.0
	expoente := 1.0

	f.Println("\n--- Tabela de Aproximação do Seno ---")
	f.Println("Termo | Expoente | Valor Acumulado")

	for i := 1; i <= 4; i++ {
		fatorial := 1.0

		for j := 1.0; j <= expoente; j++ {
			fatorial = fatorial * j
		}
		termoAtual := (m.Pow(a, expoente) / fatorial) * sinal
		
		senA = senA + termoAtual

		f.Printf("  %d   |    %.0f     |  %.6f\n", i, expoente, senA)

		sinal = sinal * -1.0
	}

	f.Printf("\nResultado final truncado: Sen(%.2f) ≈ %.6f\n", a, senA)
}
