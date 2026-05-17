package handlers

import (
	"fmt"
	"net/http"

	"sistema-saude-crud/app/utils"
)

func ListPatientsHandler(w http.ResponseWriter, r *http.Request) {
	patients, err := utils.GetPatients()
	if err != nil {
		fmt.Fprintf(w, "<h1>Erro ao listar pacientes</h1>")
		fmt.Fprintf(w, "<p>%s</p>", err.Error())
		fmt.Fprintf(w, `<br><a href="/">Voltar para o início</a>`)
		return
	}

	fmt.Fprintf(w, "<h1>Lista de Pacientes</h1>")

	if len(patients) == 0 {
		fmt.Fprintf(w, "<p>Nenhum paciente cadastrado.</p>")
		fmt.Fprintf(w, `<br><a href="/">Voltar para o início</a>`)
		return
	}

	fmt.Fprintf(w, "<table border='1' cellpadding='8'>")
	fmt.Fprintf(w, "<tr>")
	fmt.Fprintf(w, "<th>ID</th>")
	fmt.Fprintf(w, "<th>Nome</th>")
	fmt.Fprintf(w, "<th>CPF</th>")
	fmt.Fprintf(w, "<th>Email</th>")
	fmt.Fprintf(w, "<th>Telefone</th>")
	fmt.Fprintf(w, "<th>Data de nascimento</th>")
	fmt.Fprintf(w, "<th>Endereço</th>")
	fmt.Fprintf(w, "<th>Tipo sanguíneo</th>")
	fmt.Fprintf(w, "</tr>")

	for _, patient := range patients {
		fmt.Fprintf(w, "<tr>")
		fmt.Fprintf(w, "<td>%d</td>", patient.ID)
		fmt.Fprintf(w, "<td>%s</td>", patient.Nome)
		fmt.Fprintf(w, "<td>%s</td>", patient.CPF)
		fmt.Fprintf(w, "<td>%s</td>", patient.Email)
		fmt.Fprintf(w, "<td>%s</td>", patient.Telefone)
		fmt.Fprintf(w, "<td>%s</td>", patient.DataNascimento)
		fmt.Fprintf(w, "<td>%s</td>", patient.Endereco)
		fmt.Fprintf(w, "<td>%s</td>", patient.TipoSanguineo)
		fmt.Fprintf(w, "</tr>")
	}

	fmt.Fprintf(w, "</table>")
	fmt.Fprintf(w, `<br><a href="/">Voltar para o início</a>`)
}