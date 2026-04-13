package main

import f "fmt"

func main() {
	var nota1,nota2, alunos_aprovados, alunos_exame, alunos_reprovado, N int 
	var media, soma_media, med_classes float64
	
	f.Print("Digite a quantidade de alunos (N): ")
	f.Scan(&N)
	
	for i := 0; i < N; i++{	

	f.Print("Digite sua 1° nota: ")
	f.Scan(&nota1)
	f.Print("Digite sua 2° nota( Ou 0 para sair): ")
	f.Scan(&nota2)
    
		media = (float64(nota1)+float64(nota2))/2
		soma_media += media 
        f.Printf("A média do aluno foi: %.2f ",media )
	
	if media >= 7{
		alunos_aprovados++
		f.Println("Você foi aprovado. ")
	}else{
		if media < 7 && media > 3{
        alunos_exame++
		f.Println("Você está em exame. ")
	
		}else{
			if media <= 3{
				alunos_reprovado++
				f.Println("Você está reprovado. ")
			}
		}
      }
   	}
                  if N > 0{
			med_classes = soma_media/float64(N)
		  }
	
	f.Println("Total de alunos aprovados: ",alunos_aprovados)
	f.Println("Total de alunos em exame: ",alunos_exame)
	f.Println("Total de alunos reprovados: ",alunos_reprovado)
	f.Println("A media da turma é de: ",med_classes)
}
