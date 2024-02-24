package student_leave

const (
	getListRequestQuery = `
		SELECT
			slr.id,
			s.nim_number,
			s.name,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			slr.start_date,
			slr.total_leave_duration_semester,
			slr.permit_number,
			slr.purpose,
			slr.remarks,
			slr.is_approved,
			rs.semester_type,
			rs.semester_start_year
		FROM student_leave_requests slr
		JOIN students s ON s.id = slr.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters rs ON rs.id = slr.request_semester_id
	`

	countListRequestQuery = `
		SELECT COUNT(1)
		FROM student_leave_requests slr
		JOIN students s ON s.id = slr.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters rs ON rs.id = slr.request_semester_id
	`

	getDetailRequestQuery = `
		SELECT
			slr.id,
			s.nim_number,
			s.name,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			slr.start_date,
			slr.total_leave_duration_semester,
			slr.permit_number,
			slr.purpose,
			slr.remarks,
			slr.is_approved
		FROM student_leave_requests slr
		JOIN students s ON s.id = slr.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE slr.id = $1
	`

	getListQuery = `
		SELECT
			sl.id,
			s.nim_number,
			s.name,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			slv.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			se.semester_type,
			se.semester_start_year,
			slr.permit_number,
			slr.purpose,
			slr.remarks
		FROM student_leaves sl
		JOIN student_leave_requests slr ON slr.id = sl.student_leave_request_id
		JOIN students s ON s.id = slr.student_id
		JOIN semesters se ON se.id = sl.semester_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels slv ON slv.id = dsp.study_level_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM student_leaves sl
		JOIN student_leave_requests slr ON slr.id = sl.student_leave_request_id
		JOIN students s ON s.id = slr.student_id
		JOIN semesters se ON se.id = sl.semester_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels slv ON slv.id = dsp.study_level_id
	`

	getDetailQuery = `
		SELECT
			sl.id,
			s.nim_number,
			s.name,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			slv.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			se.semester_type,
			se.semester_start_year,
			slr.permit_number,
			slr.purpose,
			slr.remarks
		FROM student_leaves sl
		JOIN student_leave_requests slr ON slr.id = sl.student_leave_request_id
		JOIN students s ON s.id = slr.student_id
		JOIN semesters se ON se.id = sl.semester_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels slv ON slv.id = dsp.study_level_id
		WHERE sl.id = $1
	`

	createQuery = `
		INSERT INTO student_leave_requests (
			student_id,
			total_leave_duration_semester,
			start_date,
			permit_number,
			purpose,
			remarks
		) VALUES (
			:student_id,
			:total_leave_duration_semester,
			:start_date,
			:permit_number,
			:purpose,
			:remarks
		)
	`

	updateQuery = `
		UPDATE student_leave_requests SET
			permit_number = :permit_number,
			purpose = :purpose,
			remarks = :remarks
		WHERE id = :id
	`

	approveQuery = `
		UPDATE student_leave_requests SET is_approved = $2
		WHERE id = $1
	`

	endQuery = `
		UPDATE student_leave_requests SET
			current_leave_duration_semester = total_leave_duration_semester
		WHERE id = $1
	`

	deleteQuery = `
		DELETE FROM student_leave_requests WHERE id = $1
	`
)
