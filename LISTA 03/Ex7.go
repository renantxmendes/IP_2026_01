package main

import f "fmt"

func main() {
	
	var num, soma_num, qnt_num, maior, menor,qnt_pares, soma_pares, qnt_impares int
	var media, med_pares, porcentagem float64
	
    f.Print("Digite um número inteiro (ou 30000 para encerrar): ")
	f.Scan(&num)

	for num != 30000{
	
 	qnt_num++
	soma_num = soma_num + num 
	
	if qnt_num == 1 {
		maior = num
		menor = num
	}else{
	  if num > maior{
		maior = num 
	  }else{
		if num < menor{
			menor = num 
		}
	  }
	}
	if num % 2 == 0{
    qnt_pares++
    soma_pares = soma_pares + num 

	}else{
		qnt_impares++
	}
	
    f.Print("Digite um número inteiro (ou 30000 para encerrar): ")
	f.Scan(&num)
    }
	    media = float64(soma_num)/ float64(qnt_num)
		med_pares = float64(soma_pares) / float64(qnt_pares)
       porcentagem = (float64(qnt_impares)/float64(qnt_num)) * 100

	    f.Println("A soma dos números é: ", soma_num)
		 f.Println("A quantidade de números é: ", qnt_num )
		 f.Println("O maior dos números é: ", maior )
		  f.Println("O menor dos números é: ", menor )
	        f.Println("A média dos números é: ", media )
	          f.Println("A média dos números pares é: ", med_pares)
	            f.Printf("A porcentagem de números ímpares é:", porcentagem, "%")

				if qnt_pares > 0 {
            med_pares = float64(soma_pares) / float64(qnt_pares)
            f.Println("f) A média dos números pares é:", med_pares)
        } else {
            f.Println("f) Nenhum número par foi digitado.")
        }
        
    
        f.Println("Você encerrou o programa sem digitar nenhum número válido.")
    }
 
