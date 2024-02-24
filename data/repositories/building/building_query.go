package building

const (
	getListQuery = `
		SELECT 
			b.id,
			f.id AS faculty_id,
			f.name AS faculty_name,
			m.id AS major_id,
			m.name AS major_name,
			mf.id AS major_faculty_id,
			mf.name AS major_faculty_name,
			b.code,
			b.name
		FROM buildings b
		LEFT JOIN faculties f ON f.id = b.faculty_id
		LEFT JOIN majors m ON m.id = b.major_id
		LEFT JOIN faculties mf ON mf.id = m.faculty_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM buildings b
	`

	getDetailQuery = `
		SELECT 
			b.id,
			f.id AS faculty_id,
			f.name AS faculty_name,
			m.id AS major_id,
			m.name AS major_name,
			b.code,
			b.name
		FROM buildings b
		LEFT JOIN faculties f ON f.id = b.faculty_id
		LEFT JOIN majors m ON m.id = b.major_id
		WHERE b.id = $1
	`

	createQuery = `
		INSERT INTO buildings (
			faculty_id,
			major_id,
			code,
			name,
			created_by
		) VALUES ($1, $2, $3, $4, $5)
	`

	updateQuery = `
		UPDATE buildings SET
			faculty_id = $1,
			major_id = $2,
			code = $3,
			name = $4,
			updated_by = $5
		WHERE id = $6
	`

	deleteQuery = `
		DELETE FROM buildings WHERE id = $1
	`
)
