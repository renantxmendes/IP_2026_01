package main
import f "fmt"
func main() {
	var id, idMaior, idMenor int
	var peso, maiorPeso, menorPeso float64

	for i := 1; i <= 90; i++ {
		f.Printf("Boi %d - Digite o ID e o peso: ", i)
		f.Scan(&id, &peso)

		if i == 1 {
			idMaior = id
			maiorPeso = peso
			idMenor = id
			menorPeso = peso
		} else {
			if peso > maiorPeso {
				maiorPeso = peso
				idMaior = id
			}
			if peso < menorPeso {
				menorPeso = peso
				idMenor = id
			}
		}
	}

	f.Printf("\nBoi mais gordo -> ID: %d | Peso: %.2f\n", idMaior, maiorPeso)
	f.Printf("Boi mais magro -> ID: %d | Peso: %.2f\n", idMenor, menorPeso)
}
