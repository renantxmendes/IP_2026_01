package handlers

import (
	"fmt"
	"net/http"

	"sistema-saude-crud/app/utils"
)

func DeletePatientHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")

	err := utils.DeletePatient(id)
	if err != nil {
		fmt.Fprintf(w, "<h1>Erro ao excluir paciente</h1>")
		fmt.Fprintf(w, "<p>%s</p>", err.Error())
		fmt.Fprintf(w, `<br><a href="/forms/deletePatient.html">Voltar</a>`)
		return
	}

	fmt.Fprintf(w, "<h1>Paciente excluído com sucesso!</h1>")
	fmt.Fprintf(w, "<p><strong>ID excluído:</strong> %s</p>", id)
	fmt.Fprintf(w, `<br><a href="/listPatients">Ver lista de pacientes</a>`)
	fmt.Fprintf(w, `<br><br><a href="/">Voltar para o início</a>`)
}