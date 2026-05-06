package main

import f "fmt"

func main() {
	var notas [15] int
	var frequenciaAbsoluta [11] int

	
	f.Println("Digite as 15 notas (entre 0 e 10):")
	for i := 0; i < 15; i++ {
		for {
			f.Printf("Nota do %dº aluno: ", i+1)
			f.Scan(&notas[i])

			if notas[i] >= 0 && notas[i] <= 10 {
				break
			}
			f.Println("Nota inválida! Digite um valor entre 0 e 10.")
		}
		
		frequenciaAbsoluta[notas[i]]++
	}


	f.Println("\n-------------------------------------------------")
	f.Printf("%-10s | %-18s | %-18s\n", "Nota", "Freq. Absoluta", "Freq. Relativa")
	f.Println("-------------------------------------------------")

	for nota := 0; nota <= 10; nota++ {
		fa := frequenciaAbsoluta[nota]
		
		fr := float64(fa) / 15.0

		f.Printf("Nota %-5d | %-18d | %-18.2f\n", nota, fa, fr)
	}
	f.Println("-------------------------------------------------")
}