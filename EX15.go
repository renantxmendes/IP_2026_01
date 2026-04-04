package main

import  "fmt"
var num int

func main() {
	fmt.Print("Digite um número entre 5 e 2000: ")
	fmt.Scan(&num)
	for i := 2; i <= num ; i = i + 2 {
		fmt.Printf("%d ^ %d = %d\n", i, i, i*i)
	}
}
