package main
import f "fmt"
func main() {
	var val float64 
	var ar_cond, pintura, vidro, dir_hidra string 
	f.Print("Qual é o valor inicial do carro: ")
	f.Scan(&val)
	f.Print("Gostaria de adicionar ar condicionado? (S ou N): ")
	f.Scan(&ar_cond)
	f.Print("Gostaria de adicionar pintura metálica? (S ou N): ")
	f.Scan(&pintura)
	f.Print("Gostaria de adicionar vidro elétrico? (S ou N): ")
	f.Scan(&vidro)
    f.Print("Gostaria de adicionar direção hidráulica? (S ou N): ")
	f.Scan(&dir_hidra)
	if ar_cond == "S"{
		val = val + 1750
	}
    if pintura == "S"{
		val = val + 800
	}
    if vidro == "S"{
		val = val + 1200
	}
    if dir_hidra == "S"{
	   val = val + 2000
	}
	f.Printf("O valor do carro será de: %.2f", +val)
}
