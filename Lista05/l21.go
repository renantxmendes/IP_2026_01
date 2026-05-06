package main
import f "fmt"
func main() {
	var cod int
	var vet [10]float64
	f.Println("Digite o código (0, 1 ou 2):")
	f.Scan(&cod)
	if cod == 0 {
		return
	}
	for i := 0; i < 10; i++ {
		f.Printf("Vetor[%d]: ", i)
		f.Scan(&vet[i])
	}
	if cod == 1 {
		f.Println("Ordem direta:", vet)
	} else if cod == 2 {
		f.Print("Ordem inversa: [ ")
		for i := 9; i >= 0; i-- {
			f.Printf("%.2f ", vet[i])
		}
		f.Println("]")
	}
}