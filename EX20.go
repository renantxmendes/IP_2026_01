package main

import "fmt"

var horas, min, seg, seg_totais int

func main() {
	fmt.Print("Digite as horas, minutos e segundos: ")
	fmt.Scan(&horas, &min, &seg)

	seg_totais = horas*3600 + min*60 + seg

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d", seg_totais)
}
