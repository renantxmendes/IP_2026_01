package main

import (
	"fmt"
	"sort"
)


type Empregado struct {
	ID    int
	Meses int
}

func main() {
	var empregados []Empregado

	fmt.Println("Digite o ID do empregado e o número de meses (0 0 para encerrar):")

	for i := 0; i < 100; i++ {
		var id, meses int
		fmt.Printf("Empregado %d: ", i+1)
		fmt.Scan(&id, &meses)

	
		if id == 0 && meses == 0 {
			break
		}

		
		empregados = append(empregados, Empregado{ID: id, Meses: meses})
	}

	sort.Slice(empregados, func(i, j int) bool {
		return empregados[i].Meses < empregados[j].Meses
	})

	fmt.Println("\n--- Os 3 Empregados Mais Recentes ---")
	
	limite := 3
	if len(empregados) < 3 {
		limite = len(empregados)
	}

	if limite == 0 {
		fmt.Println("Nenhum empregado foi cadastrado.")
	} else {
		for i := 0; i < limite; i++ {
			fmt.Printf("%dº Mais Recente: ID %d (%d meses de trabalho)\n", 
				i+1, empregados[i].ID, empregados[i].Meses)
		}
	}
}