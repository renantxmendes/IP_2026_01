package main
import f "fmt"
func main() {
	var A , B , C int
	var maior , menor , inter int
	f.Print("Digite três números inteiros diferentes entre sí: ")
	f.Scan(&A,&B,&C)
	
	if A < B && A < C{
		menor = A
		if B < C{
			inter = B
			maior = C
		}else{
			inter = C
			maior = B
		}
	}else
	if B < A && B < C {
		menor = B
		if A < C{
			inter = A 
			maior = C
		}else{
			inter = C
			maior = A
		}
	}else
	if C < A && C < B {
		menor = C
		if B < A{
			inter = B
			maior = A 
		}else{
			inter = A 
			maior = B
		}
	}
        f.Println(menor , inter, maior)
}
