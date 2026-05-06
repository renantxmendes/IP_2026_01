package main
import f "fmt"
func main() {
	var janela [24]int
	var corredor [24]int
	var op, lugar int
	for {
		f.Println("\n1.Janela 2.Corredor 3.Sair")
		f.Scan(&op)
		if op == 3 { break }
		v := &janela
		txt := "Janela"
		if op == 2 { 
			v = &corredor
			txt = "Corredor"
		}
		cheio := true
		f.Printf("Poltronas Livres (%s): ", txt)
		for i := 0; i < 24; i++ {
			if v[i] == 0 {
				f.Printf("%d ", i+1)
				cheio = false
			}
		}
		if cheio {
			f.Println("Setor Lotado!")
			continue
		}
		f.Print("\nEscolha a poltrona: ")
		f.Scan(&lugar)
		if lugar < 1 || lugar > 24 || v[lugar-1] == 1 {
			f.Println("Venda Inválida!")
		} else {
			v[lugar-1] = 1
			f.Println("Venda Realizada!")
		}
	}
}