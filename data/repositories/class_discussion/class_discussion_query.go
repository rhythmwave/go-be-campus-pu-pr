package class_discussion

const (
	getListQuery = `
		SELECT
			cd.id,
			cd.class_id,
			cd.title,
			cd.abstraction,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cd.total_comment,
			cd.last_comment
		FROM class_discussions cd
		JOIN lecturers l ON l.id = cd.lecturer_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM class_discussions cd
		JOIN lecturers l ON l.id = cd.lecturer_id
	`

	getDetailQuery = `
		SELECT
			cd.id,
			cd.class_id,
			cd.title,
			cd.abstraction,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cd.total_comment,
			cd.last_comment
		FROM class_discussions cd
		JOIN lecturers l ON l.id = cd.lecturer_id
		WHERE cd.id = $1
	`

	getCommentQuery = `
		SELECT
			cdc.id,
			s.id AS student_id,
			s.nim_number AS student_nim_number,
			s.name AS student_name,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cdc.comment
		FROM class_discussion_comments cdc
		LEFT JOIN students s ON s.id = cdc.student_id
		LEFT JOIN lecturers l ON l.id = cdc.lecturer_id
	`

	countCommentQuery = `
		SELECT COUNT(1)
		FROM class_discussion_comments cdc
		LEFT JOIN students s ON s.id = cdc.student_id
		LEFT JOIN lecturers l ON l.id = cdc.lecturer_id
	`

	getDetailCommentQuery = `
	SELECT
		cdc.id,
		s.id AS student_id,
		s.nim_number AS student_nim_number,
		s.name AS student_name,
		l.id AS lecturer_id,
		l.name AS lecturer_name,
		l.front_title AS lecturer_front_title,
		l.back_degree AS lecturer_back_degree,
		cdc.comment
	FROM class_discussion_comments cdc
	LEFT JOIN students s ON s.id = cdc.student_id
	LEFT JOIN lecturers l ON l.id = cdc.lecturer_id
	WHERE cdc.id = $1
`

	createQuery = `
		INSERT INTO class_discussions (
			lecturer_id,
			class_id,
			title,
			abstraction
		) VALUES (
			:lecturer_id,
			:class_id,
			:title,
			:abstraction
		)
	`

	updateQuery = `
		UPDATE class_discussions SET
			title = :title,
			abstraction = :abstraction
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM class_discussions WHERE id = $1
	`

	createCommentQuery = `
		INSERT INTO class_discussion_comments (
			class_discussion_id,
			lecturer_id,
			student_id,
			comment
		) VALUES (
			:class_discussion_id,
			:lecturer_id,
			:student_id,
			:comment
		)
	`

	deleteCommentQuery = `
		DELETE FROM class_discussion_comments WHERE id = $1
	`
)
