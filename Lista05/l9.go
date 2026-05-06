package main
import f "fmt"
func main() {
	var atletas[10] float64	
	var soma, media float64

	for i:=0; i < len(atletas); i++{
		f.Printf("Digite a altura do %d° atleta:\n ",i+1)
		f.Scan(&atletas[i])
		
		soma+= atletas[i]
	}
	media = soma / 10

	f.Printf(" A média de altura é de: %f\n",media)

	for i:=0; i < len(atletas); i++{
		if atletas[i] > media{
			f.Println(atletas[i])
		}
	}
}