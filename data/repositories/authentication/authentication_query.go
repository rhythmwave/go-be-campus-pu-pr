package authentication

const (
	getByUsernameQuery = `
		SELECT
			a.id,
			a.username,
			a.password,
			a.admin_id,
			a.lecturer_id,
			a.student_id,
			ad.role_id AS admin_role_id,
			r.name AS admin_role_name,
			ad.name AS admin_name,
			l.name AS lecturer_name,
			s.name AS student_name,
			a.is_active,
			a.suspension_remarks,
			a.id_sso,
			a.sso_refresh_token
		FROM authentications a
		LEFT JOIN admins ad ON ad.id = a.admin_id
		LEFT JOIN roles r ON r.id = ad.role_id
		LEFT JOIN lecturers l ON l.id = a.lecturer_id
		LEFT JOIN students s ON s.id = a.student_id
	`

	getByIdQuery = `
		SELECT
			a.id,
			a.username,
			a.password,
			a.admin_id,
			a.lecturer_id,
			a.student_id,
			ad.role_id AS admin_role_id,
			r.name AS admin_role_name,
			ad.name AS admin_name,
			l.name AS lecturer_name,
			s.name AS student_name,
			a.is_active,
			a.suspension_remarks,
			a.id_sso,
			a.sso_refresh_token
		FROM authentications a
		LEFT JOIN admins ad ON ad.id = a.admin_id
		LEFT JOIN roles r ON r.id = ad.role_id
		LEFT JOIN lecturers l ON l.id = a.lecturer_id
		LEFT JOIN students s ON s.id = a.student_id
		WHERE a.id = $1
	`

	getByUserIdQuery = `
		SELECT
			a.id,
			a.username,
			a.password,
			a.admin_id,
			a.lecturer_id,
			a.student_id,
			ad.role_id AS admin_role_id,
			r.name AS admin_role_name,
			ad.name AS admin_name,
			l.name AS lecturer_name,
			s.name AS student_name,
			a.is_active,
			a.suspension_remarks,
			a.id_sso,
			a.sso_refresh_token
		FROM authentications a
		LEFT JOIN admins ad ON ad.id = a.admin_id
		LEFT JOIN roles r ON r.id = ad.role_id
		LEFT JOIN lecturers l ON l.id = a.lecturer_id
		LEFT JOIN students s ON s.id = a.student_id
	`

	createQuery = `
		INSERT INTO authentications (
			username,
			password,
			admin_id,
			lecturer_id,
			student_id
		) VALUES (
			:username,
			:password,
			:admin_id,
			:lecturer_id,
			:student_id
		) ON CONFLICT (username) DO UPDATE SET
			password = EXCLUDED.password
	`

	updatePasswordQuery = `
		UPDATE authentications SET password = $1 WHERE admin_id = $2 OR lecturer_id = $2 OR student_id = $2
	`

	updateActivationQuery = `
		UPDATE authentications SET 
			is_active = $1,
			suspension_remarks = $2
		WHERE admin_id = $3 OR lecturer_id = $3 OR student_id = $3
	`

	deleteQuery = `
		DELETE FROM authentications WHERE admin_id = $1 OR lecturer_id = $1 OR student_id = $1
	`

	updateIdSsoQuery = `
		UPDATE authentications SET id_sso = $1
		WHERE id = $2
	`

	updateSsoRefreshTokenQuery = `
		UPDATE authentications SET sso_refresh_token = $1
		WHERE id = $2
	`
)
