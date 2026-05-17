package utils

func UpdatePatient(
	id string,
	nome string,
	cpf string,
	email string,
	telefone string,
	dataNascimento string,
	endereco string,
	tipoSanguineo string,
) error {
	query := `
		UPDATE patients
		SET 
			nome = $1,
			cpf = $2,
			email = $3,
			telefone = $4,
			data_nascimento = $5,
			endereco = $6,
			tipo_sanguineo = $7
		WHERE id = $8
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
		id,
	)

	return err
}