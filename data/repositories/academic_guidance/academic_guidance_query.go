package academic_guidance

const (
	getListStudentQuery = `
		SELECT
			s.id,
			s.nim_number,
			s.student_force,
			s.name,
			s.status
		FROM academic_guidance_students ags
		JOIN academic_guidances ag ON ag.id = ags.academic_guidance_id
		JOIN students s ON s.id = ags.student_id
	`

	countListStudentQuery = `
		SELECT COUNT(1)
		FROM academic_guidance_students ags
		JOIN academic_guidances ag ON ag.id = ags.academic_guidance_id
		JOIN students s ON s.id = ags.student_id
	`

	getDetailQuery = `
		SELECT
			ag.id,
			ag.semester_id,
			ag.lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			ag.decision_number,
			ag.decision_date,
			ag.total_student
		FROM academic_guidances ag
		JOIN lecturers l ON l.id = ag.lecturer_id
		WHERE ag.id = $1
	`

	getDetailBySemesterIdLecturerIdQuery = `
		SELECT
			ag.id,
			ag.semester_id,
			ag.lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			ag.decision_number,
			ag.decision_date,
			ag.total_student
		FROM academic_guidances ag
		JOIN lecturers l ON l.id = ag.lecturer_id
		WHERE ag.semester_id = $1 AND ag.lecturer_id = $2
	`

	getDetailBySemesterIdStudentIdQuery = `
		SELECT
			ag.id,
			ag.semester_id,
			ag.lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			ag.decision_number,
			ag.decision_date,
			ag.total_student
		FROM academic_guidances ag
		JOIN lecturers l ON l.id = ag.lecturer_id
		JOIN academic_guidance_students ags ON ags.academic_guidance_id = ag.id
		WHERE ags.semester_id = $1 AND ags.student_id = $2
	`

	upsertQuery = `
		INSERT INTO academic_guidances (
			semester_id,
			lecturer_id,
			created_by
		) VALUES ($1, $2, $3)
		ON CONFLICT (semester_id, lecturer_id) DO UPDATE SET
			updated_by = EXCLUDED.created_by
		RETURNING id
	`

	upsertStudentQuery = `
		INSERT INTO academic_guidance_students (
			academic_guidance_id,
			student_id,
			created_by
		) VALUES (
			:academic_guidance_id,
			:student_id,
			:created_by
		) ON CONFLICT (semester_id, student_id) DO UPDATE SET
			academic_guidance_id = EXCLUDED.academic_guidance_id,
			updated_by = EXCLUDED.created_by
	`

	upsertDecisionQuery = `
		INSERT INTO academic_guidances (
			lecturer_id,
			semester_id,
			decision_number,
			decision_date,
			created_by
		) VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (semester_id, lecturer_id) DO UPDATE SET
			decision_number = EXCLUDED.decision_number,
			decision_date = EXCLUDED.decision_date,
			updated_by = EXCLUDED.created_by
	`

	getSessionQuery = `
		SELECT
			ags.id,
			ags.academic_guidance_id,
			ag.semester_id,
			ag.lecturer_id,
			ags.subject,
			ags.session_date,
			ags.summary
		FROM academic_guidance_sessions ags
		JOIN academic_guidances ag ON ag.id = ags.academic_guidance_id
	`

	getSessionByIdQuery = `
	SELECT
		ags.id,
		ags.academic_guidance_id,
		ag.semester_id,
		ag.lecturer_id,
		ags.subject,
		ags.session_date,
		ags.summary
	FROM academic_guidance_sessions ags
	JOIN academic_guidances ag ON ag.id = ags.academic_guidance_id
	WHERE ags.id = $1
`

	getSessionStudentQuery = `
		SELECT
			s.id,
			agss.academic_guidance_session_id,
			s.name,
			s.nim_number
		FROM academic_guidance_session_students agss
		JOIN students s ON s.id = agss.student_id
		WHERE agss.academic_guidance_session_id IN (SELECT UNNEST($1::uuid[]))
	`

	getSessionFileQuery = `
		SELECT
			agsf.id,
			agsf.academic_guidance_session_id,
			agsf.title,
			agsf.file_path,
			agsf.file_path_type
		FROM academic_guidance_session_files agsf
		WHERE agsf.academic_guidance_session_id IN (SELECT UNNEST($1::uuid[]))
	`

	upsertSessionQuery = `
		INSERT INTO academic_guidance_sessions (
			id,
			academic_guidance_id,
			subject,
			session_date,
			summary
		) VALUES (
			:id,
			:academic_guidance_id,
			:subject,
			:session_date,
			:summary
		) ON CONFLICT (id) DO UPDATE SET
			subject = EXCLUDED.subject,
			session_date = EXCLUDED.session_date,
			summary = EXCLUDED.summary
	`

	deleteSessionStudentExcludingStudentIdsQuery = `
		DELETE FROM academic_guidance_session_students WHERE academic_guidance_session_id = $1 %s
	`

	deleteSessionFileExcludingFilePathsQuery = `
		DELETE FROM academic_guidance_session_files WHERE academic_guidance_session_id = $1 %s
	`

	upsertSessionStudentQuery = `
		INSERT INTO academic_guidance_session_students (
			academic_guidance_session_id,
			student_id
		) VALUES (
			:academic_guidance_session_id,
			:student_id
		) ON CONFLICT (academic_guidance_session_id, student_id) DO NOTHING
	`

	upsertSessionFileQuery = `
		INSERT INTO academic_guidance_session_files (
			academic_guidance_session_id,
			title,
			file_path,
			file_path_type
		) VALUES (
			:academic_guidance_session_id,
			:title,
			:file_path,
			:file_path_type
		) ON CONFLICT (file_path, file_path_type) DO UPDATE SET
			title = EXCLUDED.title
	`

	deleteSessionQuery = `
		DELETE FROM academic_guidance_sessions WHERE id = $1
	`
)
