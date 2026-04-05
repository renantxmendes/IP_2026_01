package main
import f "fmt"
var tipo_cons string
var cons_agua, preco float64

func main() {
	f.Print("Digite o tipo de consumidor e o cunsumo de água: ")
	f.Scan(&tipo_cons, &cons_agua)
	if tipo_cons == "Residencial" {
		preco = 5 + cons_agua*0.05
	} else 
	if tipo_cons == "Comercial" {
		preco = 500 + (cons_agua - 80)*0.25
	} else
	if tipo_cons == "Industrial" {
		preco = 800 + (cons_agua - 100)*0.04
	}

	f.Printf("Conta %s. Valor: %.2f", tipo_cons, preco)
}
