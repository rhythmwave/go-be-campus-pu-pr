package class_exam

const (
	getListQuery = `
		SELECT
			ce.id,
			ce.class_id,
			ce.title,
			ce.abstraction,
			ce.file_path,
			ce.file_path_type,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			ce.start_time,
			ce.end_time,
			ce.total_submission
			%s
		FROM class_exams ce
		JOIN lecturers l ON l.id = ce.lecturer_id
		%s
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM class_exams ce
		JOIN lecturers l ON l.id = ce.lecturer_id
		%s
	`

	getDetailQuery = `
		SELECT
			ce.id,
			ce.class_id,
			ce.title,
			ce.abstraction,
			ce.file_path,
			ce.file_path_type,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			ce.start_time,
			ce.end_time,
			ce.total_submission
		FROM class_exams ce
		JOIN lecturers l ON l.id = ce.lecturer_id
		WHERE ce.id = $1
	`

	getSubmissionQuery = `
		SELECT
			ces.id,
			s.id AS student_id,
			s.nim_number,
			s.name,
			sp.name AS study_program_name,
			ces.file_path,
			ces.file_path_type,
			ces.point
		FROM student_classes sc
		JOIN classes c ON c.id = sc.class_id AND c.id = $1
		JOIN students s ON s.id = sc.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN class_exam_submissions ces ON ces.student_id = s.id AND ces.study_exam_id = $2
	`

	countSubmissionQuery = `
		SELECT COUNT(1)
		FROM student_classes sc
		JOIN classes c ON c.id = sc.class_id AND c.id = $1
		JOIN students s ON s.id = sc.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN class_exam_submissions ces ON ces.student_id = s.id AND ces.study_exam_id = $2
	`

	createQuery = `
		INSERT INTO class_exams (
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
		UPDATE class_exams SET
			title = :title,
			abstraction = :abstraction,
			file_path = :file_path,
			file_path_type = :file_path_type,
			start_time = :start_time,
			end_time = :end_time
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM class_exams WHERE id IN (SELECT UNNEST($1::uuid[]))
	`

	gradeSubmissionQuery = `
		INSERT INTO class_exam_submissions (
			class_exam_id,
			student_id,
			file_path,
			file_path_type,
			point
		) VALUES (
			:class_exam_id,
			:student_id,
			:file_path,
			:file_path_type,
			:point
		) ON CONFLICT (class_exam_id, student_id) DO UPDATE SET
			point = EXCLUDED.point
	`

	submitQuery = `
		INSERT INTO class_exam_submissions (
			class_exam_id,
			student_id,
			file_path,
			file_path_type
		) VALUES ($1, $2, $3, $4) ON CONFLICT (class_exam_id, student_id) DO UPDATE SET
		file_path = EXCLUDED.file_path,
		file_path_type = EXCLUDED.file_path_type
	`
)
