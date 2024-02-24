package class_work

const (
	getListQuery = `
		SELECT
			cw.id,
			cw.class_id,
			cw.title,
			cw.abstraction,
			cw.file_path,
			cw.file_path_type,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cw.start_time,
			cw.end_time,
			cw.total_submission
			%s
		FROM class_works cw
		JOIN lecturers l ON l.id = cw.lecturer_id
		%s
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM class_works cw
		JOIN lecturers l ON l.id = cw.lecturer_id
		%s
	`

	getDetailQuery = `
		SELECT
			cw.id,
			cw.class_id,
			cw.title,
			cw.abstraction,
			cw.file_path,
			cw.file_path_type,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cw.start_time,
			cw.end_time,
			cw.total_submission
		FROM class_works cw
		JOIN lecturers l ON l.id = cw.lecturer_id
		WHERE cw.id = $1
	`

	getSubmissionQuery = `
		SELECT
			cws.id,
			s.id AS student_id,
			s.nim_number,
			s.name,
			sp.name AS study_program_name,
			cws.file_path,
			cws.file_path_type,
			cws.point
		FROM student_classes sc
		JOIN classes c ON c.id = sc.class_id AND c.id = $1
		JOIN students s ON s.id = sc.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN class_work_submissions cws ON cws.student_id = s.id AND cws.class_work_id = $2
	`

	countSubmissionQuery = `
		SELECT COUNT(1)
		FROM student_classes sc
		JOIN classes c ON c.id = sc.class_id AND c.id = $1
		JOIN students s ON s.id = sc.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN class_work_submissions cws ON cws.student_id = s.id AND cws.class_work_id = $2
	`

	createQuery = `
		INSERT INTO class_works (
			lecturer_id,
			class_id,
			title,
			abstraction,
			file_path,
			file_path_type,
			start_time,
			end_time
		) VALUES (
			:lecturer_id,
			:class_id,
			:title,
			:abstraction,
			:file_path,
			:file_path_type,
			:start_time,
			:end_time
		)
	`

	updateQuery = `
		UPDATE class_works SET
			title = :title,
			abstraction = :abstraction,
			file_path = :file_path,
			file_path_type = :file_path_type,
			start_time = :start_time,
			end_time = :end_time
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM class_works WHERE id IN (SELECT UNNEST($1::uuid[]))
	`

	gradeSubmissionQuery = `
		INSERT INTO class_work_submissions (
			class_work_id,
			student_id,
			file_path,
			file_path_type,
			point
		) VALUES (
			:class_work_id,
			:student_id,
			:file_path,
			:file_path_type,
			:point
		) ON CONFLICT (class_work_id, student_id) DO UPDATE SET
			point = EXCLUDED.point
	`

	submitQuery = `
		INSERT INTO class_work_submissions (
			class_work_id,
			student_id,
			file_path,
			file_path_type
		) VALUES ($1, $2, $3, $4) ON CONFLICT (class_work_id, student_id) DO UPDATE SET
			file_path = EXCLUDED.file_path,
			file_path_type = EXCLUDED.file_path_type
	`
)
