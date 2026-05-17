package utils

func InsertPatient(
	nome string,
	cpf string,
	email string,
	telefone string,
	dataNascimento string,
	endereco string,
	tipoSanguineo string,
) error {
	query := `
		INSERT INTO patients 
		(nome, cpf, email, telefone, data_nascimento, endereco, tipo_sanguineo)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := DB.Exec(
		query,
		nome,
		cpf,
		email,
		telefone,
		dataNascimento,
		endereco,
		tipoSanguineo,
	)

	return err
}