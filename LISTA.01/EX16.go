package main

import "fmt"

var sal, sal_reajuste float64

func main() {
	fmt.Print("Qual o salário do funcionário: ")
	fmt.Scan(&sal)

	if sal <= 300 {
		sal_reajuste= sal*1.5
	} else {
		sal_reajuste = sal*1.3
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f", sal_reajuste)
}
