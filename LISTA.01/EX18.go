package main

import  "fmt"

var val_ini, razao, qnt_term, soma int
var termos []int

func main() {
	fmt.Print("Digite o valor inicial, a razão e a quantidade de termos da P.A: ")
	fmt.Scan(&val_ini, &razao, &qnt_term)

	termos = append(termos, val_ini)
	
	for i:= 1 ; i < qnt_term ; i++ {
		termos = append(termos, termos[i - 1] + razao)
	}	
	for i := 1; i <= qnt_term ; i++ {
		soma = soma + termos[i - 1]
	}
	fmt.Print(soma)
}
