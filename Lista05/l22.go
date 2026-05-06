package main
import f "fmt"
func main() {
	var codigos [10]int
	var saldos [10]float64
	var op, conta, i int
	var valor float64
	for i = 0; i < 10; i++ {
		f.Printf("Código e saldo da conta %d: ", i+1)
		f.Scan(&codigos[i], &saldos[i])
	}
	for {
		f.Println("\n1.Depósito 2.Saque 3.Ativo 4.Sair")
		f.Scan(&op)
		if op == 4 { break }
		if op == 1 || op == 2 {
			f.Print("Código da conta: ")
			f.Scan(&conta)
			indice := -1
			for j := 0; j < 10; j++ {
				if codigos[j] == conta {
					indice = j
					break
				}
			}
			if indice == -1 {
				f.Println("Conta não encontrada")
			} else if op == 1 {
				f.Print("Valor: ")
				f.Scan(&valor)
				saldos[indice] += valor
			} else {
				f.Print("Valor: ")
				f.Scan(&valor)
				if saldos[indice] >= valor {
					saldos[indice] -= valor
				} else {
					f.Println("Saldo insuficiente")
				}
			}
		} else if op == 3 {
			total := 0.0
			for _, s := range saldos { total += s }
			f.Printf("Ativo Total: %.2f\n", total)
		}
	}
}