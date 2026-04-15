package main
import f "fmt"
func main(){

var x float64
var soma float64

f.Print("Digite o valor de x: ")
f.Scan(&x)

fatorial:= 1.0
sinal:= 1.0

f.Println("\nCalculando os 20 primeiros termos...")

for i:= 0; i <= 20; i++{
	if i > 0{
		fatorial = fatorial * float64(i) 
	   }
	   termo := sinal *  (x/fatorial)
	   soma += termo
	   sinal *= -1.0
	}
	f.Printf("O resultado do somatório S é: %4.f\n", soma)
}
