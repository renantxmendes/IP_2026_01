package main
import f "fmt"
func main() {
	var vet [10]int
	for i := 0; i < 10; i++ {
		f.Printf("Vetor[%d]: ", i)
		f.Scan(&vet[i])
	}
	for i := 0; i < 10; i++ {
		cont := 0
		for j := 1; j <= vet[i]; j++ {
			if vet[i]%j == 0 {
				cont++
			}
		}
		if cont == 2 {
			f.Printf("Número primo: %d na posição %d\n", vet[i], i)
		}
	}
}