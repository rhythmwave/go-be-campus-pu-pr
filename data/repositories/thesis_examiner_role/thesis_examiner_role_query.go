package thesis_examiner_role

const (
	getListQuery = `
		SELECT
			ter.id,
			ter.name,
			ter.remarks,
			ter.sort
		FROM thesis_examiner_roles ter
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM thesis_examiner_roles ter
	`

	getDetailQuery = `
		SELECT 
			ter.id,
			ter.name,
			ter.remarks,
			ter.sort
		FROM thesis_examiner_roles ter
		WHERE ter.id = $1
	`

	createQuery = `
		INSERT INTO thesis_examiner_roles (
			name,
			remarks,
			sort,
			created_by
		) VALUES ($1, $2, $3, $4);
	`

	updateQuery = `
		UPDATE thesis_examiner_roles SET
			name = $1,
			remarks = $2,
			sort = $3,
			updated_by = $4
		WHERE id = $5
	`

	deleteQuery = `
		DELETE FROM thesis_examiner_roles WHERE id = $1
	`
)
