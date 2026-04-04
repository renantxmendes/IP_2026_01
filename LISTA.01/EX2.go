package main

import "fmt"

var qnt_jogos int
var cat_pop, cat_ger, cat_arq, cat_cad, qnt_pes, renda float64

func main() {
	fmt.Print("Informe a quantidade de jogos: ")
	fmt.Scan(&qnt_jogos)

	var jogos []float64

	for i := 0; i < qnt_jogos; i++ {
		fmt.Println("Informe a quantidade de pessoas no jogo", i+1, "e as porcentagens em cada categoria: ")
		fmt.Scan(&qnt_pes, &cat_pop, &cat_arq, &cat_cad)
		calcularRenda()
		jogos = append(jogos, renda)
	}

	for i := 0; i < qnt_jogos; i++ {
		fmt.Printf("A RENDA DO JOGO SERÁ DE N.%d: %.2f\n", i+1, jogos[i])
	}
}
func calcularRenda() {

	var qnt_pop, qnt_ger, qnt_arq, qnt_cad float64

	qnt_pop = qnt_pes * cat_pop / 100
	qnt_ger = qnt_pes * cat_ger / 100
	qnt_arq = qnt_pes * cat_arq / 100
	qnt_cad = qnt_pes * cat_cad / 100

	renda = (qnt_pop*1 + qnt_ger*5 + qnt_arq*10 + qnt_cad*20)
}
