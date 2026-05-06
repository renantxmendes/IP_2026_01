package main
import f "fmt"
func main() {
	var vetor[10] int
	var Pares[] int
	var Impares[] int
	var soma int

	for i:=0; i < len(vetor); i++{
		f.Printf(" Digite o valor do vetor[%d] = \n",1+i)
		f.Scan(&vetor[i])
	}
	
	for i:= 0; i < len(vetor); i++{
		if vetor[i]%2==0{
			termo:= vetor[i]
			Pares = append(Pares, termo)
		
			soma+= vetor[i] 
			
		}else{
			termo:= vetor[i]
			Impares = append(Impares, termo)
			
		}
	}
	f.Printf("Esses são os números pares = %v\n",Pares)
	f.Printf("Essa é a soma dos números pares = %v\n",soma)
	f.Printf("Esses são os números ímpares = %v\n",Impares)
	f.Printf("A quantidade de números ímpares é = %v\n", len(Impares))
}