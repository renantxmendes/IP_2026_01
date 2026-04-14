package main

import f "fmt"

func main() {
	var n, prox int
	termo1 := 0
	termo2 := 1

	
	f.Print("Quantos termos haverá sua sequência: ")
	f.Scan(&n)
	
	f.Println("\n A sequência de Fibonacci:")
	 if n >=1{
		f.Println(termo1, " ")
	 }
	 if n >=2{
		f.Println(termo2," ")
	 }
    
	 for i := 3; i <= n; i++{
		prox = termo1 + termo2
		f.Println(prox," ")
      termo1 = termo2
	  termo2 = prox

	 }
     
}
