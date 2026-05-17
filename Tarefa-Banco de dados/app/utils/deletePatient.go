package utils

func DeletePatient(id string) error {
	query := `
		DELETE FROM patients
		WHERE id = $1
	`

	_, err := DB.Exec(query, id)

	return err
}