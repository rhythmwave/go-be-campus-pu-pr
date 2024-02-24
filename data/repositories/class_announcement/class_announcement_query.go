package class_announcement

const (
	getListQuery = `
		SELECT
			ca.id,
			ca.class_id,
			ca.title,
			ca.content,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			ca.file_path,
			ca.file_path_type,
			ca.start_time,
			ca.end_time
		FROM class_announcements ca
		JOIN lecturers l ON l.id = ca.lecturer_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM class_announcements ca
		JOIN lecturers l ON l.id = ca.lecturer_id
	`

	getDetailQuery = `
		SELECT
			ca.id,
			ca.class_id,
			ca.title,
			ca.content,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			ca.file_path,
			ca.file_path_type,
			ca.start_time,
			ca.end_time
		FROM class_announcements ca
		JOIN lecturers l ON l.id = ca.lecturer_id
		WHERE ca.id = $1
	`

	createQuery = `
		INSERT INTO class_announcements (
			lecturer_id,
			class_id,
			title,
			content,
			file_path,
			file_path_type,
			start_time,
			end_time
		) VALUES (
			:lecturer_id,
			:class_id,
			:title,
			:content,
			:file_path,
			:file_path_type,
			:start_time,
			:end_time
		)
	`

	updateQuery = `
		UPDATE class_announcements SET
			title = :title,
			content = :content,
			file_path = :file_path,
			file_path_type = :file_path_type,
			start_time = :start_time,
			end_time = :end_time
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM class_announcements WHERE id IN (SELECT UNNEST($1::uuid[]))
	`
)
