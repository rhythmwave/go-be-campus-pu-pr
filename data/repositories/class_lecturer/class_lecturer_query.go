package class_lecturer

const (
	getByClassIdLecturerIdQuery = `
		SELECT
			cl.id,
			cl.class_id,
			cl.lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cl.is_grading_responsible
		FROM class_lecturers cl
		JOIN lecturers l ON l.id = cl.lecturer_id
		WHERE cl.class_id = $1 AND cl.lecturer_id = $2
	`
)
