package main
import f "fmt"
func main() {
	var A[10] int
	var jaContados[10] bool		

	for i:=0; i < len(A); i++{
		f.Printf("Digite o valor do vetor x[%d] =\n ",i+1)
		f.Scan(&A[i])
	}
	f.Println(" --- Relátorio de repetições ---")
	teveRepeticao:= false 

	for i:=0; i < len(A); i++{
		if jaContados[i]{
			continue
		}
		quantidade:=1

		for j:=i+1; j < len(A); j++{
			if A[i] == A[j]{
				quantidade++
				jaContados[j] = true
			}
		}
		if quantidade > 1{
			f.Printf(" O número %d repetiu %d vezes \n",A[i], quantidade)
			teveRepeticao = true
		}
	}	
	if ! teveRepeticao{
		f.Println("Não teve repetição.")
	}



	
}