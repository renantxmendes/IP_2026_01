package main
import f "fmt"
func main() {
	var fibonacci[50] int
	fibonacci[0] = 1
	fibonacci[1] = 1	

	for i:=2; i < len(fibonacci); i++{
		
		fibonacci[i] = fibonacci[i -1 ] + fibonacci[i - 2]
	}
	f.Println(fibonacci)
	
}	