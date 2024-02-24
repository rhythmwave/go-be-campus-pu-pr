package role

const (
	getListQuery = `
		SELECT
			r.id,
			r.name,
			r.description
		FROM roles r
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM roles r
	`

	getDetailQuery = `
		SELECT
			r.id,
			r.name,
			r.description
		FROM roles r
		WHERE r.id = $1
	`

	createQuery = `
		INSERT INTO roles (
			name,
			description,
			created_by
		) VALUES ($1, $2, $3) RETURNING id;
	`

	createStudyProgramQuery = `
		INSERT INTO role_study_program (role_id, study_program_id)
		SELECT $1, sp.id
		FROM study_programs sp
		WHERE sp.id IN (SELECT UNNEST($2::uuid[]))
		ON CONFLICT (role_id, study_program_id) DO NOTHING
	`

	createPermissionQuery = `
		INSERT INTO role_permission (role_id, permission_id)
		SELECT $1, p.id
		FROM permissions p
		WHERE p.id IN (SELECT UNNEST($2::uuid[]))
		ON CONFLICT (role_id, permission_id) DO NOTHING
	`

	updateQuery = `
		UPDATE roles SET
			name = $1,
			description = $2,
			updated_by = $3
		WHERE id = $4
	`

	updateDeleteExcludedStudyProgramQuery = `
		DELETE FROM role_study_program
		WHERE role_id = $1 AND study_program_id NOT IN (SELECT UNNEST($2::uuid[]))
	`

	updateDeleteExcludedPermissionQuery = `
		DELETE FROM role_permission
		WHERE role_id = $1 AND permission_id NOT IN (SELECT UNNEST($2::uuid[]))
	`

	deleteQuery = `
		DELETE FROM roles WHERE id = $1
	`
)
