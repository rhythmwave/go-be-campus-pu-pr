package grade_type

const (
	getListQuery = `
		SELECT
			gt.id,
			sl.id AS study_level_id,
			sl.short_name AS study_level_short_name,
			gt.code,
			gt.grade_point,
			gt.minimum_grade,
			gt.maximum_grade,
			gt.grade_category,
			gt.grade_point_category,
			gt.label,
			gt.english_label,
			gt.start_date,
			gt.end_date
		FROM grade_types gt
		JOIN study_levels sl ON sl.id = gt.study_level_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM grade_types gt
		JOIN study_levels sl ON sl.id = gt.study_level_id
	`

	getDetailQuery = `
		SELECT
			gt.id,
			sl.id AS study_level_id,
			sl.short_name AS study_level_short_name,
			gt.code,
			gt.grade_point,
			gt.minimum_grade,
			gt.maximum_grade,
			gt.grade_category,
			gt.grade_point_category,
			gt.label,
			gt.english_label,
			gt.start_date,
			gt.end_date
		FROM grade_types gt
		JOIN study_levels sl ON sl.id = gt.study_level_id
		WHERE gt.id = $1
	`

	createQuery = `
		INSERT INTO grade_types (
			study_level_id,
			code,
			grade_point,
			minimum_grade,
			maximum_grade,
			grade_category,
			grade_point_category,
			label,
			english_label,
			start_date,
			end_date,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);
	`

	updateQuery = `
		UPDATE grade_types SET
			code = $1,
			grade_point = $2,
			minimum_grade = $3,
			maximum_grade = $4,
			grade_category = $5,
			grade_point_category = $6,
			label = $7,
			english_label = $8,
			start_date = $9,
			end_date = $10,
			updated_by = $11
		WHERE id = $12
	`

	deleteQuery = `
		DELETE FROM grade_types WHERE id = $1
	`

	getByGradeCodeQuery = `
		SELECT
			gt.id,
			sl.id AS study_level_id,
			sl.short_name AS study_level_short_name,
			gt.code,
			gt.grade_point,
			gt.minimum_grade,
			gt.maximum_grade,
			gt.grade_category,
			gt.grade_point_category,
			gt.label,
			gt.english_label,
			gt.start_date,
			gt.end_date
		FROM grade_types gt
		JOIN study_levels sl ON sl.id = gt.study_level_id
		WHERE gt.study_level_id = $1 AND gt.code = $2
	`
)
