package main
import f "fmt"
func main() {
	var a, b int
	f.Print("Digite dois números inteiros: ")
	f.Scan(&a,&b)
	if a%b == 0{
		f.Println("O número",a,"é divisível por",b)
	}else{
        f.Println("O número",a,"não é divisível por",b)
	}
	
}
	
		
	
