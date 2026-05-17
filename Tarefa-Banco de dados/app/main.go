package main

import (
	"fmt"
	"log"
	"net/http"

	"sistema-saude-crud/app/handlers"
	"sistema-saude-crud/app/utils"
)

func main() {
	utils.ConnectToDB()

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)

	http.HandleFunc("/createPatient", handlers.CreatePatientHandler)
	http.HandleFunc("/listPatients", handlers.ListPatientsHandler)
	http.HandleFunc("/updatePatient", handlers.UpdatePatientHandler)
	http.HandleFunc("/deletePatient", handlers.DeletePatientHandler)

	fmt.Println("Servidor rodando em: http://localhost:3000")

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}