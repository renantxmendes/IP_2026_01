package main
import f "fmt"
func somaArray(arr []float64) float64 {
	if len(arr) == 0 {
		return 0.0
	}
	return arr[0] + somaArray(arr[1:])
}

func main() {
	valores := []float64{2.5, 3.5, 4.0, 10.0}

	resultado := somaArray(valores)

	f.Printf("Soma dos valores: %.2f\n", resultado)
}
