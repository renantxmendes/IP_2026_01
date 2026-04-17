package main
import f "fmt"
func main(){
	var soma float64
	n1:=37
	n2:=38

	for n3:=1; n3 <=37; n3++{
		termo:= float64(n1*n2)/float64(n3)
		n1--
		n2--
		soma+= termo
	}
	f.Printf("A soma será de:%f ",soma)
}
