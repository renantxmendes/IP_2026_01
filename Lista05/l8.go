package main
import f "fmt"
import m "math"
func main() {
	var x[15] float64
	var raiz[15]float64
//Preenchendo vetores
	for i:=0; i < len(x); i++{
		f.Printf("Digite o valor do vetor x [%d] = ",1+i)
		f.Scan(&x[i])
	}
	f.Println("\nExibindo a raiz quadrada do vetor x")

	for i:=0; i < len(x); i++{
		if x[i] >= 0 {
			raiz[i] = m.Sqrt(x[i])
			
		}else{
			raiz[i] = -1
		}
		f.Printf("Os valores do vetor raiz [%d] = %1.f\n", 1+i, raiz[i])
	}
	
		
}