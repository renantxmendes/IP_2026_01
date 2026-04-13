package main

import f "fmt"

func main() {
	// qnt_pes = quantidade de pessoas + 50. qnt_med = quantidade média . por_peso = porcentagem de peso
	var idade, qnt_idosos, qnt_jovens, total_pes, pes_leves int 
	var altura, peso, soma_alturas, porcentagem float64
	continuar := 1

    for continuar == 1 {
	
	f.Print("Digite sua idade: ")
	f.Scan(&idade)

	f.Print("Digite sua altura: ")
    f.Scan(&altura)

	f.Print("Digite o seu peso: ")
	f.Scan(&peso)
   
	total_pes++
	
	if idade > 50{
			qnt_idosos++ 
	}
	if idade >= 10 && idade <= 20{
    qnt_jovens++
	soma_alturas = altura + soma_alturas
	}			
	if peso < 40{
		pes_leves++
	}
	f.Print("Você gostaria de adicionar outra pessoa( 1 - Sim ,  2- Não): ")
	f.Scan(&continuar )
	}
	
	f.Println("Total de pessoas com mais de 50 anos: ", qnt_idosos )
	if qnt_jovens > 0{
		var media float64 
		media = soma_alturas / float64(qnt_jovens)
		f.Println("A media das alturas dos jovens é de: ", media)
	}else{
		f.Println("Nenhuma pessoa entre 10 e 20 anos foi cadastrada.")
	}
    if total_pes > 0{
		
		porcentagem = (float64(qnt_jovens) / float64(total_pes)) * 100
		f.Println("A porcentagem de pessoas com menos de 40Kg: ", porcentagem)
	}
}
