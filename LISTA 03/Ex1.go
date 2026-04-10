package main

import f "fmt"

func main() {
	var base, expoente int 
	resultado := 1
	
	f.Print("Digite o número da base: ")
	f.Scan(&base)
	
	f.Println("Digite o número do expoente: ")
	f.Scan(&expoente)
	
	for i := 1; i <= expoente; i++{
		resultado = resultado * base

		
	}
	
    f.Println(" O resultado da potência é igual a : ",resultado)
}
