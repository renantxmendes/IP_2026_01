package main

import "fmt"

var sal_min, qnt_kw float64		


func main() {
	fmt.Print("Escreva o valor do salário mínimo: ")
	fmt.Scan(&sal_min)
	fmt.Print("Escreva a quantidade de Kw gastas: ")
	fmt.Scan(&qnt_kw)

	cust_kw := (sal_min*0.7)/100
	cust_cons := cust_kw * qnt_kw
	desconto := cust_cons * 0.9

	fmt.Printf("Custo do kilowatt é de: R$ %.2f\n ", cust_kw)
    fmt.Printf("Custo total: R$ %.2f\n", cust_cons)
	fmt.Printf("O valor com desconto foi de: R$ %.2f\n", desconto)

}
	
