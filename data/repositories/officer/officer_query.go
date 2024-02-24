package officer

const (
	getListQuery = `
		SELECT 
			o.id,
			o.id_national_lecturer,
			o.name,
			o.title,
			o.english_title,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			o.signature_path,
			o.signature_path_type,
			o.employee_no,
			o.show_signature
		FROM officers o
		LEFT JOIN study_programs sp ON sp.id = o.study_program_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM officers o
		LEFT JOIN study_programs sp ON sp.id = o.study_program_id
	`

	getDetailQuery = `
		SELECT 
			o.id,
			o.id_national_lecturer,
			o.name,
			o.title,
			o.english_title,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			o.signature_path,
			o.signature_path_type,
			o.employee_no,
			o.show_signature
		FROM officers o
		LEFT JOIN study_programs sp ON sp.id = o.study_program_id
		WHERE o.id = $1
	`

	createQuery = `
		INSERT INTO officers (
			id_national_lecturer,
			name,
			title,
			english_title,
			study_program_id,
			signature_path,
			signature_path_type,
			show_signature,
			employee_no,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`

	updateQuery = `
		UPDATE officers SET
			id_national_lecturer = $1,
			name = $2,
			title = $3,
			english_title = $4,
			study_program_id = $5,
			signature_path = $6,
			signature_path_type = $7,
			show_signature = $8,
			employee_no = $9,
			updated_by = $10
		WHERE id = $11
	`

	deleteQuery = `
		DELETE FROM officers WHERE id = $1
	`
)
