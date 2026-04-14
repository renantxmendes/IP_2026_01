package main

import f "fmt"

func main() {
	var num int
	fatorial := 1

	f.Print("Digite um número inteiro para calcular seu fatorial: ")
	f.Scan(&num)
	if num < 0{
		f.Println("Número inválido.")
	} else if num == 0{
		f.Println("O fatorial de 0 é igual a 1")

	}else{
		for i := 1; i <= num; i++{
			fatorial = fatorial * i
		}
		f.Printf("O fatorial de %d é: %d\n", num, fatorial)
	}

}
