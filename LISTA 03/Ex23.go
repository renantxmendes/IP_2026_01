package main
import f "fmt"
func main() {
	var n int
	num:=1000.0
	den:=1.0
	sinal:=1.0
	soma:=0.0
	f.Print("Digite o número de termos: ")
	f.Scan(&n)

	if n <=0{
		f.Print("Número inváido. ")
		return
	}
	for i:=1; i <=n; i++ {
		termo_atual:= (num/den) * sinal 
		soma= soma + termo_atual	
		
		num = num - 3.0
		den = den + 1.0
		sinal = sinal * - 1.0
	}
	f.Printf("\nO resultado da série com %d termos é: %.4f\n", n, soma)

}
