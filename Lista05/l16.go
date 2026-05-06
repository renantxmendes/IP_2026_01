package main

import f "fmt"

func main() {
	var idades [50]int
	
	frequencia := make(map[int]int)

	f.Println("Digite as 50 idades:")

	for i := 0; i < 50; i++ {
		f.Printf("%dª idade: ", i+1)
		f.Scan(&idades[i])

		
		frequencia[idades[i]]++
	}
	maiorFrequencia := 0
	moda := 0

	for idade, contagem := range frequencia {
		if contagem > maiorFrequencia {
			maiorFrequencia = contagem
			moda = idade
		}
	}

	f.Println("\n------------------------------")
	if maiorFrequencia > 1 {
		f.Printf("A moda das idades é: %d\n", moda)
		f.Printf("Ela apareceu %d vezes.\n", maiorFrequencia)
	} else {
		f.Println("Não existe moda (todas as idades aparecem apenas uma vez).")
	}
	f.Println("------------------------------")
}