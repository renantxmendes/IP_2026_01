package main

import f "fmt"

func main() {

	var numero int 
 
	f.Print("Digite um número inteiro (ou 0 para sair): ")
	f.Scan(&numero)	

	for numero > 0{
		
		var testador int = 1 
		var eQuadrado bool = false 
		
		for testador * testador <= numero {

		if testador * testador == numero {
			eQuadrado = true 

		}
        testador++
		}
	   if eQuadrado == true{
		f.Print("O número é um quadrado perfeito")
	   }else{
		f.Println(" O número não é um quadrado perfeito")
	   }
	   f.Println("Digite um número inteiro (ou 0 para sair): ")
	   f.Scan(&numero)	
	}

}

