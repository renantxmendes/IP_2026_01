package main
import f "fmt"
func main() {
	var num [10]int
	var divis [5]int
	for i := 0; i < 10; i++ {
		f.Printf("Num[%d]: ", i)
		f.Scan(&num[i])
	}
	for i := 0; i < 5; i++ {
		f.Printf("Divis[%d]: ", i)
		f.Scan(&divis[i])
	}
	for i := 0; i < 10; i++ {
		f.Printf("\nNúmero %d:", num[i])
		encontrou := false
		for j := 0; j < 5; j++ {
			if num[i]%divis[j] == 0 {
				f.Printf("\n  Divisível por %d na posição %d", divis[j], j)
				encontrou = true
			}
		}
		if !encontrou {
			f.Printf("\n  Não possui divisores no segundo vetor")
		}
		f.Println()
	}
}