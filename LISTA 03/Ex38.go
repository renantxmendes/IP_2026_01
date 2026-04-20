package main
import f "fmt"
func main() {
	var cpfBase string

	f.Print("Digite os 9 primeiros digitos do CPF (sem pontuacao): ")
	f.Scan(&cpfBase)

	soma := 0
	multiplicador := 2

	for i := 8; i >= 0; i-- {
		digito := int(cpfBase[i] - '0')
		soma = soma + (digito * multiplicador)
		multiplicador++
	}
	resto := soma % 11
	digitoVerificador := 11 - resto

	if digitoVerificador >= 10 {
		digitoVerificador = 0
	}
	f.Printf("Primeiro digito verificador calculado: %d\n", digitoVerificador)
}
