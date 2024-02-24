package class_material

const (
	getListQuery = `
		SELECT
			cm.id,
			cm.title,
			cm.abstraction,
			cm.file_path,
			cm.file_path_type,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cm.is_active,
			cm.created_at
		FROM class_materials cm
		JOIN lecturers l ON l.id = cm.lecturer_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM class_materials cm
		JOIN lecturers l ON l.id = cm.lecturer_id
	`

	getDetailQuery = `
		SELECT
			cm.id,
			cm.title,
			cm.abstraction,
			cm.file_path,
			cm.file_path_type,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cm.is_active,
			cm.created_at
		FROM class_materials cm
		JOIN lecturers l ON l.id = cm.lecturer_id
		WHERE cm.id = $1
	`

	createQuery = `
		INSERT INTO class_materials (
			lecturer_id,
			class_id,
			title,
			abstraction,
			file_path,
			file_path_type,
			is_active
		) VALUES (
			:lecturer_id,
			:class_id,
			:title,
			:abstraction,
			:file_path,
			:file_path_type,
			:is_active
		)
	`

	updateQuery = `
		UPDATE class_materials SET
			title = :title,
			abstraction = :abstraction,
			file_path = :file_path,
			file_path_type = :file_path_type,
			is_active = :is_active
		WHERE id = :id
	`

	bulkUpdateActivationQuery = `
		UPDATE class_materials SET is_active = $2
		WHERE id IN (SELECT UNNEST($1::uuid[]))
	`

	bulkDeleteQuery = `
		DELETE FROM class_materials
		WHERE id IN (SELECT UNNEST($1::uuid[]))
	`
)
