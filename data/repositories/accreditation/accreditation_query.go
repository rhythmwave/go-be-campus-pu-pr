package accreditation

const (
	getListQuery = `
		SELECT 
			a.id,
			a.study_program_id,
			a.decree_number,
			a.decree_date,
			a.decree_due_date,
			a.accreditation,
			a.is_active
		FROM accreditations a
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM accreditations a
	`

	getDetailQuery = `
		SELECT 
			a.id,
			a.decree_number,
			a.decree_date,
			a.decree_due_date,
			a.accreditation,
			a.is_active
		FROM accreditations a
		WHERE a.id = $1
	`

	createQuery = `
		INSERT INTO accreditations (
			study_program_id,
			decree_number,
			decree_date,
			decree_due_date,
			accreditation,
			is_active,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	updateQuery = `
		UPDATE accreditations SET
			decree_number = $1,
			decree_date = $2,
			decree_due_date = $3,
			accreditation = $4,
			is_active = $5,
			updated_by = $6
		WHERE id = $7
	`

	deleteQuery = `
		DELETE FROM accreditations WHERE id = $1
	`
)
