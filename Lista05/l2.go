package main
import f "fmt"
func main() {
	var x[10] int
 	var y [5] int
	var vetorPares [] int
	var vetorImpares [] int
	var soma int
//Preenchendo os vetores 
	for i:=0; i < len(x); i++{
		f.Printf("Digite o valor de x[%d]: ", 1+i)
		f.Scan(&x[i])
	}
	for i:=0; i < len(y); i++{
		f.Printf("Digite o valor de y[%d]: ", 1+i)
		f.Scan(&y[i])
			soma = soma + y[i]

	}

	for i:=0; i < len(x); i++{
		if x[i]%2 == 0{
			termo := soma + x[i]
			vetorPares = append(vetorPares, termo)
	}else{
		termo:= x[i] + soma
		vetorImpares = append(vetorImpares, termo)
	}
  }
  f.Printf("\nPrimeiro vetor resultante (Pares): %v\n", vetorPares)
	f.Printf("Segundo vetor resultante (Ímpares): %v\n", vetorImpares)
}