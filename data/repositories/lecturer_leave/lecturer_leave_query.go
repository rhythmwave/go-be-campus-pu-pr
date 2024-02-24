package lecturer_leave

const (
	getListQuery = `
		SELECT
			ll.id,
			l.name,
			l.id_national_lecturer,
			l.front_title,
			l.back_degree,
			s.semester_start_year,
			s.semester_type,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			ll.start_date,
			ll.end_date,
			ll.permit_number,
			ll.purpose,
			ll.remarks,
			ll.is_active,
			ll.file_path,
			ll.file_path_type
		FROM lecturer_leaves ll
		JOIN lecturers l ON l.id = ll.lecturer_id
		JOIN semesters s ON s.id = ll.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM lecturer_leaves ll
		JOIN lecturers l ON l.id = ll.lecturer_id
		JOIN semesters s ON s.id = ll.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
	`

	getDetailQuery = `
		SELECT
			ll.id,
			l.name,
			l.id_national_lecturer,
			l.front_title,
			l.back_degree,
			s.semester_start_year,
			s.semester_type,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			ll.start_date,
			ll.end_date,
			ll.permit_number,
			ll.purpose,
			ll.remarks,
			ll.is_active,
			ll.file_path,
			ll.file_path_type
		FROM lecturer_leaves ll
		JOIN lecturers l ON l.id = ll.lecturer_id
		JOIN semesters s ON s.id = ll.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE ll.id = $1
	`

	getDetailByLecturerIdAndStartDateQuery = `
		SELECT
			ll.id,
			l.name,
			l.id_national_lecturer,
			l.front_title,
			l.back_degree,
			s.semester_start_year,
			s.semester_type,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			ll.start_date,
			ll.end_date,
			ll.permit_number,
			ll.purpose,
			ll.remarks,
			ll.is_active,
			ll.file_path,
			ll.file_path_type
		FROM lecturer_leaves ll
		JOIN lecturers l ON l.id = ll.lecturer_id
		JOIN semesters s ON s.id = ll.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE ll.lecturer_id = $1 AND ll.start_date = $2
	`

	createQuery = `
		INSERT INTO lecturer_leaves (
			lecturer_id,
			semester_id,
			start_date,
			permit_number,
			purpose,
			remarks,
			file_path,
			file_path_type,
			created_by
		) VALUES (
			:lecturer_id,
			:semester_id,
			:start_date,
			:permit_number,
			:purpose,
			:remarks,
			:file_path,
			:file_path_type,
			:created_by
		)
	`

	updateQuery = `
		UPDATE lecturer_leaves SET
			permit_number = :permit_number,
			purpose = :purpose,
			remarks = :remarks,
			file_path = :file_path,
			file_path_type = :file_path_type,
			updated_by = :updated_by
		WHERE id = :id
	`

	endQuery = `
		UPDATE lecturer_leaves SET
			end_date = NOW(),
			is_active = false,
			updated_by = $2
		WHERE id = $1
	`

	deleteQuery = `
		DELETE FROM lecturer_leaves WHERE id = $1
	`

	autoSetActiveQuery = `
		UPDATE lecturer_leaves SET is_active = true
		WHERE start_date = DATE(now()) RETURNING lecturer_id
	`
)
