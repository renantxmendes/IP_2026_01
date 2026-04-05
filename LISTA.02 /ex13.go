package main
import f "fmt"
func main() {
	var num, centenas, dezenas int 
	f.Print("Digite um número inteiro de 3 casas: ")
	f.Scan(&num)
	if num < 100{
		f.Print("O número é inválido: ")
		return 
	}
	centenas = num / 100
	dezenas = (num - centenas*100) / 10
	f.Println(dezenas)


	
}
