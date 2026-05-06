package main

import "fmt"

func main() {
	var v1 [10]int
	var v2 [10]int
	var resultante [20]int

	fmt.Println("Digite os 10 elementos do Vetor 1:")
	for i := 0; i < 10; i++ {
		fmt.Printf("V1[%d]: ", i)
		fmt.Scan(&v1[i])
	}

	
	fmt.Println("\nDigite os 10 elementos do Vetor 2:")
	for i := 0; i < 10; i++ {
		fmt.Printf("V2[%d]: ", i)
		fmt.Scan(&v2[i])
	}

	
	for i := 0; i < 10; i++ {
		resultante[i*2] = v1[i]     
		resultante[i*2+1] = v2[i]   
	}
	fmt.Println("\n--- Vetores Originais ---")
	fmt.Println("V1:", v1)
	fmt.Println("V2:", v2)

	fmt.Println("\n--- Vetor Resultante (Intercalado) ---")
	fmt.Print("[ ")
	for _, valor := range resultante {
		fmt.Printf("%d ", valor)
	}
	fmt.Println("]")
}