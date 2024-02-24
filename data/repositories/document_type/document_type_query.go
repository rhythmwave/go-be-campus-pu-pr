package document_type

const (
	getListQuery = `
		SELECT
			dt.id,
			dt.name
		FROM document_types dt
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM document_types dt
	`

	getDetailQuery = `
		SELECT 
			dt.id,
			dt.name
		FROM document_types dt
		WHERE dt.id = $1
	`

	createQuery = `
		INSERT INTO document_types (
			name,
			created_by
		) VALUES ($1, $2);
	`

	updateQuery = `
		UPDATE document_types SET
			name = $1,
			updated_by = $2
		WHERE id = $3
	`

	deleteQuery = `
		DELETE FROM document_types WHERE id = $1
	`
)
