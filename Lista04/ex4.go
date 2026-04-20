package main
import f "fmt"
func decimalParaBinario(n int) string {
	if n < 2 {
		return f.Sprintf("%d", n)
	}
	return decimalParaBinario(n/2) + f.Sprintf("%d", n%2)
}
func main() {
	var numero int

	f.Print("Digite um numero decimal: ")
	f.Scan(&numero)

	binario := decimalParaBinario(numero)

	f.Printf("Forma binaria: %s\n", binario)
}
