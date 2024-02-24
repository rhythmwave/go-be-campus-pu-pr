package class_grade_component

const (
	getByClassIdQuery = `
		SELECT
			cgc.id,
			cgc.class_id,
			cgc.name,
			cgc.percentage,
			cgc.is_active
		FROM class_grade_components cgc
		WHERE cgc.class_id = $1
		ORDER BY cgc.name
	`

	upsertQuery = `
		INSERT INTO class_grade_components (
			class_id,
			name,
			percentage,
			is_active
		) VALUES (
			:class_id,
			:name,
			:percentage,
			:is_active
		) ON CONFLICT (class_id, name) DO UPDATE SET
			percentage = EXCLUDED.percentage,
			is_active = EXCLUDED.is_active
	`

	deleteAllByClassIdExcludingNamesQuery = `
		DELETE FROM class_grade_components WHERE class_id = $1 AND name NOT IN (SELECT UNNEST($2::text[]))
	`

	deleteAllByClassIdQuery = `
		DELETE FROM class_grade_components WHERE class_id = $1
	`
)
