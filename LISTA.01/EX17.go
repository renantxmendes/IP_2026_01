package main

import  "fmt"

var num_par, qnt_num int
var num_pares []int

func main() {
	fmt.Print("Digite o número par inicial e quantos deverão ser mostrados: ")
	fmt.Scan(&num_par, &qnt_num)

	if num_par%2 != 0 {
		fmt.Print("O primeiro número nâo é par")
		return
	}

	for i := 0 ; i < qnt_num ; i++ {
		num_pares = append(num_pares, num_par)
		num_par = num_par + 2
	}
	fmt.Println(num_pares)
}package main

import  "fmt"

var num_par, qnt_num int
var num_pares []int

func main() {
	fmt.Print("Digite o número par inicial e quantos deverão ser mostrados: ")
	fmt.Scan(&num_par, &qnt_num)

	if num_par%2 != 0 {
		fmt.Print("O primeiro número nâo é par")
		return
	}

	for i := 0 ; i < qnt_num ; i++ {
		num_pares = append(num_pares, num_par)
		num_par = num_par + 2
	}
	fmt.Println(num_pares)
}
