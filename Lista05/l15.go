package main

import f "fmt"

func main() {
    var v1 [30]int
    var resultante [30]int

    
    for i := 0; i < 30; i++ {
        f.Printf("Digite o valor de v1[%d] = ", i+1)
        f.Scan(&v1[i])
    }

    
    for i := 0; i < 30; i++ {
        if i % 2 == 0 {
            
            resultante[i] = v1[i] * 2
        } else {
            
            resultante[i] = v1[i] * 3
        }
    }
    f.Printf("\nVetor original: %v\n", v1)
    f.Println("Vetor resultante:")
    for i, valor := range resultante {
        f.Printf("Posição %d: %d\n", i, valor)
    }
}