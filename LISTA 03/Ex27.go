package main
import f "fmt"
import m "math"
func main() {
	var x float64
	
	f.Print("Digite o valor de x (em radianos): ")
	f.Scan(&x)

	soma := 0.0
	sinal := 1.0
	expoente := 0.0
	fat := 1.0

	for i := 1; i <= 20; i++ {
		
		termoAtual := (m.Pow(x, expoente) / fat) * sinal
		
		soma = soma + termoAtual

		sinal = sinal * -1.0
		
		proxExpoente := expoente + 2.0
		
		fat = fat * proxExpoente * (proxExpoente - 1.0)
		
		expoente = proxExpoente
	}

	valorReal := m.Cos(x)
	diferenca := m.Abs(valorReal - soma)

	f.Printf("\n--- Resultados ---\n")
	

	f.Printf("a) Cosseno calculado (Série com 20 termos): %.15f\n", soma)
	f.Printf("   Valor da função math.Cos(x) do Go:       %.15f\n", valorReal)
	f.Printf("b) Diferença (Erro da aproximação):         %.15f\n", diferenca)
}
