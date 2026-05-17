package handlers

import (
	"fmt"
	"net/http"

	"sistema-saude-crud/app/utils"
)

func CreatePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	nome := r.FormValue("nome")
	cpf := r.FormValue("cpf")
	email := r.FormValue("email")
	telefone := r.FormValue("telefone")
	dataNascimento := r.FormValue("data_nascimento")
	endereco := r.FormValue("endereco")
	tipoSanguineo := r.FormValue("tipo_sanguineo")

	err := utils.InsertPatient(
		nome,
		cpf,
		email,
		telefone,
		dataNascimento,
		endereco,
		tipoSanguineo,
	)

	if err != nil {
		fmt.Fprintf(w, "<h1>Erro ao cadastrar paciente</h1>")
		fmt.Fprintf(w, "<p>%s</p>", err.Error())
		fmt.Fprintf(w, `<br><a href="/forms/createPatient.html">Voltar</a>`)
		return
	}

	fmt.Fprintf(w, "<h1>Paciente cadastrado com sucesso!</h1>")
	fmt.Fprintf(w, "<p><strong>Nome:</strong> %s</p>", nome)
	fmt.Fprintf(w, "<p><strong>CPF:</strong> %s</p>", cpf)
	fmt.Fprintf(w, "<p><strong>Email:</strong> %s</p>", email)
	fmt.Fprintf(w, "<p><strong>Telefone:</strong> %s</p>", telefone)
	fmt.Fprintf(w, "<p><strong>Data de nascimento:</strong> %s</p>", dataNascimento)
	fmt.Fprintf(w, "<p><strong>Endereço:</strong> %s</p>", endereco)
	fmt.Fprintf(w, "<p><strong>Tipo sanguíneo:</strong> %s</p>", tipoSanguineo)
	fmt.Fprintf(w, `<br><a href="/">Voltar para o início</a>`)
}