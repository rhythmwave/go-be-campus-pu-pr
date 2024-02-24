package lecturer_resignation

const (
	getListQuery = `
		SELECT
			lr.id,
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
			lr.resign_date,
			lr.resignation_number,
			lr.purpose,
			lr.remarks
		FROM lecturer_resignations lr
		JOIN lecturers l ON l.id = lr.lecturer_id
		JOIN semesters s ON s.id = lr.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM lecturer_resignations lr
		JOIN lecturers l ON l.id = lr.lecturer_id
		JOIN semesters s ON s.id = lr.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
	`

	getDetailQuery = `
		SELECT
			lr.id,
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
			lr.resign_date,
			lr.resignation_number,
			lr.purpose,
			lr.remarks
		FROM lecturer_resignations lr
		JOIN lecturers l ON l.id = lr.lecturer_id
		JOIN semesters s ON s.id = lr.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE lr.id = $1
	`

	createQuery = `
		INSERT INTO lecturer_resignations (
			lecturer_id,
			semester_id,
			resign_date,
			resignation_number,
			purpose,
			remarks,
			created_by
		) VALUES (
			:lecturer_id,
			:semester_id,
			:resign_date,
			:resignation_number,
			:purpose,
			:remarks,
			:created_by
		)
	`
)
