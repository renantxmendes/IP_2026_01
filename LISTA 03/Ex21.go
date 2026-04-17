package main
import f "fmt"
func main() {
	var b,n int
	f.Print("Digite o número de sua base: ")
	f.Scan(&b)
	f.Print("Digite o número de seu expoente: ")
	f.Scan(&n)

	resultado:=1

	for i:=0; i<n;i++{
		resultado = resultado * b
      
	}
	f.Printf("O resultado de %d elevado a %d é: %d\n", b, n, resultado)
}
	
