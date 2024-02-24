package grade_component

const (
	getListQuery = `
		SELECT
			gc.id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			sc.id AS subject_category_id,
			sc.name AS subject_category_name,
			gc.name,
			gc.is_active,
			gc.default_percentage
		FROM grade_components gc
		JOIN study_programs sp ON sp.id = gc.study_program_id
		JOIN subject_categories sc ON sc.id = gc.subject_category_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM grade_components gc
		JOIN study_programs sp ON sp.id = gc.study_program_id
		JOIN subject_categories sc ON sc.id = gc.subject_category_id
	`

	getDetailQuery = `
		SELECT
			gc.id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			sc.id AS subject_category_id,
			sc.name AS subject_category_name,
			gc.name,
			gc.is_active
		FROM grade_components gc
		JOIN study_programs sp ON sp.id = gc.study_program_id
		JOIN subject_categories sc ON sc.id = gc.subject_category_id
		WHERE gc.id = $1
	`

	createQuery = `
		INSERT INTO grade_components (
			study_program_id,
			subject_category_id,
			name,
			is_active,
			created_by
		) VALUES ($1, $2, $3, $4, $5);
	`

	updateQuery = `
		UPDATE grade_components SET
			subject_category_id = $1,
			name = $2,
			is_active = $3,
			updated_by = $4
		WHERE id = $5
	`

	deleteQuery = `
		DELETE FROM grade_components WHERE id = $1
	`

	getDistinctSubjectCategoryListQuery = `
		SELECT
			DISTINCT(sc.id) AS subject_category_id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			sc.name AS subject_category_name
		FROM grade_components gc
		JOIN study_programs sp ON sp.id = gc.study_program_id
		JOIN subject_categories sc ON sc.id = gc.subject_category_id
	`

	countDistinctSubjectCategoryListQuery = `
		WITH data AS (
			SELECT
				DISTINCT(sc.id) AS subject_category_id,
				sp.id AS study_program_id,
				sp.name AS study_program_name,
				sc.name AS subject_category_name
			FROM grade_components gc
			JOIN study_programs sp ON sp.id = gc.study_program_id
			JOIN subject_categories sc ON sc.id = gc.subject_category_id
			%s
		) SELECT COUNT(1) FROM data
	`

	getPercentageBySubjectCategoriesQuery = `
		SELECT
			gc.id,
			gc.subject_category_id,
			gc.name,
			gc.default_percentage,
			gc.is_active
		FROM grade_components gc
		WHERE gc.study_program_id = $1 AND gc.subject_category_id IN (SELECT UNNEST($2::uuid[]))
		ORDER BY gc.name
	`

	bulkUpdatePercentageQuery = `
		INSERT INTO grade_components (
			id,
			study_program_id,
			subject_category_id,
			name,
			default_percentage,
			is_active,
			created_by
		) VALUES (
			:id,
			:study_program_id,
			:subject_category_id,
			:name,
			:default_percentage,
			:is_active,
			:updated_by
		) ON CONFLICT (id) DO UPDATE SET
			default_percentage = EXCLUDED.default_percentage,
			is_active = EXCLUDED.is_active,
			updated_by = EXCLUDED.created_by
	`
)
