package main

import f "fmt"

var dia_sem, categoria string
var preco_i, preco_f float64

func main() {
	f.Print("Digite o dia da semana: ")
	f.Scan(&dia_sem)
	f.Print("Digite o preco e a categoria do DVD: ")
	f.Scan(&preco_i, &categoria)
	switch dia_sem {
	case "Segunda", "Terça", "Quinta":
		preco_f = preco_i*0.6
		if categoria == "Lançamento" {
			preco_f = preco_f*1.15
		}
	case "Quarta", "Sexta", "Sábado", "Domingo":
		preco_f = preco_i
		if categoria == "Lançamento" {
			preco_f = preco_f * 1.15
		}
	}

	f.Printf("O valor do aluguel é: %.2f", preco_f)
}
