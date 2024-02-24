package study_plan

const (
	bulkCreateQuery = `
		INSERT INTO study_plans (
			id,
			student_id,
			semester_id,
			semester_package,
			maximum_credit,
			is_thesis,
			is_submitted
		) VALUES (
			:id,
			:student_id,
			:semester_id,
			:semester_package,
			:maximum_credit,
			:is_thesis,
			:is_submitted
		) ON CONFLICT(id) DO NOTHING
	`

	bulkApproveQuery = `
		UPDATE study_plans SET is_approved = $1
		WHERE id IN (SELECT UNNEST($2::uuid[])) AND is_submitted IS true
	`

	getListQuery = `
		SELECT
			spl.id,
			spl.is_submitted,
			spl.is_approved,
			spl.semester_id,
			se.semester_start_year,
			se.semester_type,
			spl.total_mandatory_credit,
			spl.total_optional_credit,
			spl.grade_point,
			s.id AS student_id,
			s.nim_number AS student_nim_number,
			s.name AS student_name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.code AS dikti_study_program_code,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			spl.is_thesis
		FROM study_plans spl
		JOIN students s ON s.id = spl.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters se ON se.id = spl.semester_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM study_plans spl
		JOIN students s ON s.id = spl.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters se ON se.id = spl.semester_id
	`

	getApprovedByStudentIdQuery = `
		SELECT
			spl.id,
			spl.is_submitted,
			spl.is_approved,
			spl.semester_id,
			se.semester_start_year,
			se.semester_type,
			spl.total_mandatory_credit,
			spl.total_optional_credit,
			spl.grade_point,
			s.id AS student_id,
			s.nim_number AS student_nim_number,
			s.name AS student_name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.code AS dikti_study_program_code,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			spl.is_thesis
		FROM study_plans spl
		JOIN students s ON s.id = spl.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters se ON se.id = spl.semester_id
		WHERE spl.student_id = $1 AND spl.is_approved IS true
	`

	getByStudentIdAndSemesterIdQuery = `
		SELECT
			spl.id,
			spl.is_submitted,
			spl.is_approved,
			spl.semester_id,
			se.semester_start_year,
			se.semester_type,
			spl.total_mandatory_credit,
			spl.total_optional_credit,
			spl.grade_point,
			s.id AS student_id,
			s.nim_number AS student_nim_number,
			s.name AS student_name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.code AS dikti_study_program_code,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			spl.is_thesis
		FROM study_plans spl
		JOIN students s ON s.id = spl.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters se ON se.id = spl.semester_id
		WHERE spl.student_id = $1 AND spl.semester_id = $2
	`

	getByStudentIdsAndSemesterIdQuery = `
		SELECT
			spl.id,
			spl.is_submitted,
			spl.is_approved,
			spl.semester_id,
			se.semester_start_year,
			se.semester_type,
			spl.total_mandatory_credit,
			spl.total_optional_credit,
			spl.grade_point,
			s.id AS student_id,
			s.nim_number AS student_nim_number,
			s.name AS student_name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.code AS dikti_study_program_code,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			spl.is_thesis
		FROM study_plans spl
		JOIN students s ON s.id = spl.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters se ON se.id = spl.semester_id
		WHERE spl.student_id IN (SELECT UNNEST($1::uuid[])) AND spl.semester_id = $2
	`
)
