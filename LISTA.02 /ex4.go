package main

import (
	f "fmt"
	m "math"
	
)
func main() {
	var num, res float64
	f.Print("Digite um número real qualquer: ")
	f.Scan(&num)
	if num >= 0{ 
		 res = m.Sqrt(num) 
		f.Printf("O resultado será: %.2f\n ",res )
	}else{
		
			res =  (num)*(num)
			f.Printf("O resultado será: %.2f\n ", res)
		
		
	}
}
