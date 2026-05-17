package utils

type Patient struct {
	ID             int
	Nome           string
	CPF            string
	Email          string
	Telefone       string
	DataNascimento string
	Endereco       string
	TipoSanguineo  string
	CreatedAt      string
}

func GetPatients() ([]Patient, error) {
	rows, err := DB.Query(`
		SELECT 
			id, 
			nome, 
			cpf, 
			email, 
			telefone, 
			data_nascimento, 
			endereco, 
			tipo_sanguineo, 
			created_at
		FROM patients
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []Patient

	for rows.Next() {
		var patient Patient

		err := rows.Scan(
			&patient.ID,
			&patient.Nome,
			&patient.CPF,
			&patient.Email,
			&patient.Telefone,
			&patient.DataNascimento,
			&patient.Endereco,
			&patient.TipoSanguineo,
			&patient.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		patients = append(patients, patient)
	}

	return patients, nil
}