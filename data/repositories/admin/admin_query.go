package admin

const (
	getListQuery = `
		SELECT 
			a.id, 
			a.name, 
			a.username,
			r.id AS role_id,
			r.name AS role_name
		FROM admins a
		JOIN roles r ON r.id = a.role_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM admins a
		JOIN roles r ON r.id = a.role_id
	`

	getDetailQuery = `
	SELECT 
		a.id, 
		a.name, 
		a.username,
		r.id AS role_id,
		r.name AS role_name
	FROM admins a
	JOIN roles r ON r.id = a.role_id
	WHERE a.id = $1
`

	createQuery = `
		INSERT INTO admins (
			username,
			name,
			password,
			role_id,
			created_by
		) VALUES ($1, $2, $3, $4, $5)
	`

	updateQuery = `
		UPDATE admins SET
			username = $1,
			name = $2,
			password = COALESCE($3, password),
			role_id = $4,
			updated_by = $5
		WHERE id = $6
	`

	deleteQuery = `
		DELETE FROM admins WHERE id = $1
	`

	getSingleSuperAdminQuery = `
		SELECT id, name, username
		FROM admins 
		WHERE role_id IS NULL LIMIT 1;
	`
)
