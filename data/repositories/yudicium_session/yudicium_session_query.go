package yudicium_session

const (
	getListQuery = `
		SELECT
			ys.id,
			ys.semester_id,
			s.semester_start_year,
			s.semester_type,
			ys.name,
			ys.session_date
		FROM yudicium_sessions ys
		JOIN semesters s ON s.id = ys.semester_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM yudicium_sessions ys
		JOIN semesters s ON s.id = ys.semester_id
	`

	getDetailQuery = `
		SELECT 
			ys.id,
			ys.semester_id,
			s.semester_start_year,
			s.semester_type,
			ys.name,
			ys.session_date
		FROM yudicium_sessions ys
		JOIN semesters s ON s.id = ys.semester_id
		WHERE ys.id = $1
	`

	getUpcomingQuery = `
		SELECT 
			ys.id,
			ys.semester_id,
			s.semester_start_year,
			s.semester_type,
			ys.name,
			ys.session_date
		FROM yudicium_sessions ys
		JOIN semesters s ON s.id = ys.semester_id
		WHERE DATE(ys.session_date) > DATE(now())
		LIMIT 1
	`

	createQuery = `
		INSERT INTO yudicium_sessions (
			semester_id,
			name,
			session_date,
			created_by
		) VALUES ($1, $2, $3, $4);
	`

	updateQuery = `
		UPDATE yudicium_sessions SET
			semester_id = $1,
			name = $2,
			session_date = $3,
			updated_by = $4
		WHERE id = $5
	`

	deleteQuery = `
		DELETE FROM yudicium_sessions WHERE id = $1
	`

	doQuery = `
		UPDATE yudicium_sessions SET
			yudicium_number = :yudicium_number,
			actual_date = :actual_date
		WHERE id = :id
	`
)
