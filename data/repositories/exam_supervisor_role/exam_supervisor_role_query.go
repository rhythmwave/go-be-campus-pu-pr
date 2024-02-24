package exam_supervisor_role

const (
	getListQuery = `
		SELECT
			esr.id,
			esr.name,
			esr.sort
		FROM exam_supervisor_roles esr
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM exam_supervisor_roles esr
	`

	getDetailQuery = `
		SELECT 
			esr.id,
			esr.name,
			esr.sort
		FROM exam_supervisor_roles esr
		WHERE esr.id = $1
	`

	createQuery = `
		INSERT INTO exam_supervisor_roles (
			name,
			sort,
			created_by
		) VALUES ($1, $2, $3);
	`

	updateQuery = `
		UPDATE exam_supervisor_roles SET
			name = $1,
			sort = $2,
			updated_by = $3
		WHERE id = $4
	`

	deleteQuery = `
		DELETE FROM exam_supervisor_roles WHERE id = $1
	`
)
