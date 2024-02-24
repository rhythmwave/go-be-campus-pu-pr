package yudicium_term

const (
	getListQuery = `
		SELECT
			yt.id,
			c.id AS curriculum_id,
			c.name AS curriculum_name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			yt.term,
			yt.remarks
		FROM yudicium_terms yt
		JOIN curriculums c ON c.id = yt.curriculum_id
		JOIN study_programs sp ON sp.id = c.study_program_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM yudicium_terms yt
		JOIN curriculums c ON c.id = yt.curriculum_id
		JOIN study_programs sp ON sp.id = c.study_program_id
	`

	getDetailQuery = `
		SELECT 
			yt.id,
			c.id AS curriculum_id,
			c.name AS curriculum_name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			yt.term,
			yt.remarks
		FROM yudicium_terms yt
		JOIN curriculums c ON c.id = yt.curriculum_id
		JOIN study_programs sp ON sp.id = c.study_program_id
		WHERE yt.id = $1
	`

	createQuery = `
		INSERT INTO yudicium_terms (
			curriculum_id,
			term,
			remarks,
			created_by
		) VALUES ($1, $2, $3, $4);
	`

	updateQuery = `
		UPDATE yudicium_terms SET
			term = $1,
			remarks = $2,
			updated_by = $3
		WHERE id = $4
	`

	deleteQuery = `
		DELETE FROM yudicium_terms WHERE id = $1
	`
)
