package main
import f "fmt"
func main() {
	var n1, n2 int

	f.Print("Digite N1: ")
	f.Scan(&n1)
	f.Print("Digite N2: ")
	f.Scan(&n2)

	maior := n1
	if n2 > n1 {
		maior = n2
	}

	mmc := maior

	for {
		if mmc%n1 == 0 && mmc%n2 == 0 {
			break
		}
		mmc++
	}

	f.Printf("O MMC de %d e %d e: %d\n", n1, n2, mmc)
}
