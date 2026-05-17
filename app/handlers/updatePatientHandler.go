package handlers

import (
	"fmt"
	"net/http"

	"sistema-saude-crud/app/utils"
)

func UpdatePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	nome := r.FormValue("nome")
	cpf := r.FormValue("cpf")
	email := r.FormValue("email")
	telefone := r.FormValue("telefone")
	dataNascimento := r.FormValue("data_nascimento")
	endereco := r.FormValue("endereco")
	tipoSanguineo := r.FormValue("tipo_sanguineo")

	err := utils.UpdatePatient(
		id,
		nome,
		cpf,
		email,
		telefone,
		dataNascimento,
		endereco,
		tipoSanguineo,
	)

	if err != nil {
		fmt.Fprintf(w, "<h1>Erro ao atualizar paciente</h1>")
		fmt.Fprintf(w, "<p>%s</p>", err.Error())
		fmt.Fprintf(w, `<br><a href="/forms/updatePatient.html">Voltar</a>`)
		return
	}

	fmt.Fprintf(w, "<h1>Paciente atualizado com sucesso!</h1>")
	fmt.Fprintf(w, "<p><strong>ID:</strong> %s</p>", id)
	fmt.Fprintf(w, "<p><strong>Nome:</strong> %s</p>", nome)
	fmt.Fprintf(w, "<p><strong>CPF:</strong> %s</p>", cpf)
	fmt.Fprintf(w, `<br><a href="/listPatients">Ver lista de pacientes</a>`)
	fmt.Fprintf(w, `<br><br><a href="/">Voltar para o início</a>`)
}