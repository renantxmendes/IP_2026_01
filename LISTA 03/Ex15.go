package main
import f "fmt"
func main() {
	var n int
	var num int
	num = 1
	f.Print("Digite o número de termos da sua sequência: ")
	f.Scan(&n)
	for i:= 1; i<= n; i++{
    termo:= num * num
	num++
	f.Println(termo)
	}
    f.Println(" ")
}
