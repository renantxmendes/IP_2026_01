package main
import f "fmt"
func main() {
	var x[10] int 
	encontrou:= false

	for i:=0; i < len(x); i++{
		f.Printf("Digite o %d° número do vetor: ",1+i)
		f.Scan(&x[i])
		if x[i] > 50{
			f.Printf("O número %d está na posição %d\n.",x[i], 1+i)
			encontrou = true
			
		}
			
		}
		if encontrou == false{
			f.Print("Não há números acima de 50.")
			}
	}
	
