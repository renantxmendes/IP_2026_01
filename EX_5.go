package main

import "fmt"

var conta int
var cons_agua, preco float64
var local string

func main() {
	fmt.Print("Informe seu tipo de local, consumo de água e o número da conta: ")
	fmt.Scan(&local,&cons_agua,&conta)
	if local == "R"{
		preco = 5 + cons_agua * 0.05
	} else
	if local == "C"{
		preco = 500 + (cons_agua - 80) * 0.25
	} else 
	if local == "I"{
		preco = 800 + (cons_agua - 100) * 0.04
	}
	fmt.Printf( " O número da  conta será de: %d\nVALOR DA CONTA = %.2f", conta, preco)

}
