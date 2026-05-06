package main
import f "fmt"
func main() {
	var vet [10]int
	f.Println("Digite 10 números em ordem crescente:")
	for i := 0; i < 10; i++ {
		f.Printf("%dº número: ", i+1)
		f.Scan(&vet[i])
	}
	f.Println("Vetor:", vet)
}