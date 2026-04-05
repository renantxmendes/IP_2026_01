package main
import f "fmt" 
var num int

func main(){
	f.Print("Digite um número inteiro qualquer: ")
	f.Scan(&num)
	if num%2 == 0 {
		f.Printf("O número %d é PAR.\n",num)

	}else{
		f.Printf("O número %d é ÍMPAR.\n",num)
	}
		
} 
