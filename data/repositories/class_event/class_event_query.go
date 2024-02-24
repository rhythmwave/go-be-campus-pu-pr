package class_event

const (
	getListQuery = `
		SELECT
			ce.id,
			ce.title,
			ce.frequency,
			ce.event_time,
			ce.lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			ce.remarks,
			ce.is_active,
			ce.created_at
		FROM class_events ce
		JOIN lecturers l ON l.id = ce.lecturer_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM class_events ce
		JOIN lecturers l ON l.id = ce.lecturer_id
	`

	getDetailQuery = `
		SELECT
			ce.id,
			ce.title,
			ce.frequency,
			ce.event_time,
			ce.lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			ce.remarks,
			ce.is_active,
			ce.created_at
		FROM class_events ce
		JOIN lecturers l ON l.id = ce.lecturer_id
		WHERE ce.id = $1
	`

	createQuery = `
		INSERT INTO class_events (
			lecturer_id,
			class_id,
			title,
			frequency,
			event_time,
			remarks,
			is_active
		) VALUES (
			:lecturer_id,
			:class_id,
			:title,
			:frequency,
			:event_time,
			:remarks,
			:is_active
		)
	`

	updateQuery = `
		UPDATE class_events SET
			title = :title,
			frequency = :frequency,
			event_time = :event_time,
			remarks = :remarks,
			is_active = :is_active
		WHERE id = :id
	`

	bulkUpdateActivationQuery = `
		UPDATE class_events SET is_active = $2
		WHERE id IN (SELECT UNNEST($1::uuid[]))
	`

	bulkDeleteQuery = `
		DELETE FROM class_events
		WHERE id IN (SELECT UNNEST($1::uuid[]))
	`
)
