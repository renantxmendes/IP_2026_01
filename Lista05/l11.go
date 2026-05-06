package main
import f "fmt"
func main() {
var x[100] float64
var S float64


for i:=0; i < len(x); i++{
	f.Printf("Digite o valor do array x[%d] = \n",1+i)
	f.Scan(&x[i])
}

for i:=0; i < 50; i++{
	
	oposto:= 99 - i
	diferenca := x[i] - x[oposto]
	cubo:= diferenca * diferenca * diferenca

	S += cubo
}
f.Printf("\nO valor do somatório S é: %.2f\n", S)
}