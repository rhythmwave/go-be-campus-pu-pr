package yudicium_student

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
			s.total_credit,
			s.gpa,
			s.status,
			ys.application_date,
			ys.done_yudicium
		FROM yudicium_students ys
		JOIN students s ON s.id = ys.student_id
		LEFT JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		LEFT JOIN yudicium_sessions yse ON yse.id = ys.yudicium_session_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM yudicium_students ys
		JOIN students s ON s.id = ys.student_id
		LEFT JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		LEFT JOIN yudicium_sessions yse ON yse.id = ys.yudicium_session_id
	`

	createQuery = `
		INSERT INTO yudicium_students (
			student_id,
			application_date,
			with_thesis
		) VALUES (
			:student_id,
			:application_date,
			:with_thesis
		) ON CONFLICT (student_id) DO NOTHING
	`

	doQuery = `
		UPDATE yudicium_students SET yudicium_session_id = $1
		WHERE student_id IN (SELECT UNNEST($2::uuid[]))
	`
)
