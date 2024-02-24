package thesis_supervisor_role

const (
	getListQuery = `
		SELECT
			tsr.id,
			tsr.name,
			tsr.sort
		FROM thesis_supervisor_roles tsr
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM thesis_supervisor_roles tsr
	`

	getDetailQuery = `
		SELECT 
			tsr.id,
			tsr.name,
			tsr.sort
		FROM thesis_supervisor_roles tsr
		WHERE tsr.id = $1
	`

	getFirstOrderQuery = `
		SELECT 
			tsr.id,
			tsr.name,
			tsr.sort
		FROM thesis_supervisor_roles tsr
		ORDER BY tsr.sort ASC LIMIT 1
	`

	createQuery = `
		INSERT INTO thesis_supervisor_roles (
			name,
			sort,
			created_by
		) VALUES ($1, $2, $3);
	`

	updateQuery = `
		UPDATE thesis_supervisor_roles SET
			name = $1,
			sort = $2,
			updated_by = $3
		WHERE id = $4
	`

	deleteQuery = `
		DELETE FROM thesis_supervisor_roles WHERE id = $1
	`
)
