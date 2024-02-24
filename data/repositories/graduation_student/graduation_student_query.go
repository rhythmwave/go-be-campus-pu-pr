package graduation_student

const (
	getListQuery = `
		SELECT
			s.id,
			s.nim_number,
			s.name,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			gs.application_date
		FROM graduation_students gs
		JOIN students s ON s.id = gs.student_id
		LEFT JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM graduation_students gs
		JOIN students s ON s.id = gs.student_id
		LEFT JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
	`

	createQuery = `
		INSERT INTO graduation_students (
			student_id,
			application_date,
			graduation_session_id
		) VALUES (
			:student_id,
			:application_date,
			:graduation_session_id
		) ON CONFLICT (student_id) DO NOTHING
	`
)
