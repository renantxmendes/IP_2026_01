package main

import f "fmt"

func main() {
	
	var soma, media, quantidade int
	soma = 0 
	quantidade = 0

	for i:= 50; i <= 70; i++{
		if i%2 == 0{
			soma = soma + i
			quantidade++
			
			media = soma/ quantidade
			
		}
	}
	f.Println("A soma dos pares será de: ",soma)
	f.Println("A media será de: ", media)
	
}
