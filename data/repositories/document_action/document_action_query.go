package document_action

const (
	getListQuery = `
		SELECT
			da.id,
			da.action,
			da.english_action
		FROM document_actions da
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM document_actions da
	`

	getDetailQuery = `
		SELECT 
			da.id,
			da.action,
			da.english_action
		FROM document_actions da
		WHERE da.id = $1
	`

	createQuery = `
		INSERT INTO document_actions (
			action,
			english_action,
			created_by
		) VALUES ($1, $2, $3);
	`

	updateQuery = `
		UPDATE document_actions SET
			action = $1,
			english_action = $2,
			updated_by = $3
		WHERE id = $4
	`

	deleteQuery = `
		DELETE FROM document_actions WHERE id = $1
	`
)
