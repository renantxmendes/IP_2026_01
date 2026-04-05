 package main 
 import f "fmt"
 
 func main (){
var num int 
	
    f.Print("Digite um número inteiro qualquer: ")
	f.Scan(&num)
    if num < 90 && num > 20{
		f.Println("O número",num,"está dentro da faixa")
	}else{
		f.Println("O número",num,"não está dentro da faixa")
	}
	
 }
