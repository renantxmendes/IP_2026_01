package main
import f "fmt"
func main() {
	var x[10] int

	for i:=0; i < len(x); i++{
		f.Printf( "Digite o valor do vetor x[%d] = \n",1+i)
		f.Scan(&x[i])
	}
	f.Println("--- Análise do menor valor do vetor x ---")

	menor:= x[0]
	posicao:= 0

	for i:=1; i < len(x); i++{
		if x[i] < menor{
			menor = x[i]
			posicao = i
		}
	}
	f.Printf("O menor valor será %d e ele está na posição %d", menor, posicao)
}