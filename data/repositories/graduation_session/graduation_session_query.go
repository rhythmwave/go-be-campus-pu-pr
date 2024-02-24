package graduation_session

const (
	getListQuery = `
		SELECT
    gs.id,
    gs.session_year,
    gs.session_number,
    gs.session_date
		FROM graduation_sessions gs
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM graduation_sessions gs
	`

	getDetailQuery = `
		SELECT 
			gs.id,
			gs.session_year,
			gs.session_number,
			gs.session_date
		FROM graduation_sessions gs
		WHERE gs.id = $1
	`

	getUpcomingQuery = `
		SELECT 
			gs.id,
			gs.session_year,
			gs.session_number,
			gs.session_date
		FROM graduation_sessions gs
		WHERE DATE(gs.session_date) > DATE(now())
		LIMIT 1
	`

	createQuery = `
		INSERT INTO graduation_sessions (
			session_year,
			session_number,
			session_date,
			created_by
		) VALUES (
			:session_year,
			:session_number,
			:session_date,
			:created_by
		);
	`

	updateQuery = `
		UPDATE graduation_sessions SET
    session_year = :session_year,
    session_number = :session_number,
    session_date = :session_date,
    updated_by = :updated_by
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM graduation_sessions WHERE id = $1
	`
)
