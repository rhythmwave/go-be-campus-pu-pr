package subject_grade_component

const (
	getBySubjectIdQuery = `
		SELECT
			sgc.id,
			sgc.subject_id,
			sgc.name,
			sgc.percentage,
			sgc.is_active
		FROM subject_grade_components sgc
		WHERE sgc.subject_id = $1
		ORDER BY sgc.name
	`

	upsertQuery = `
		INSERT INTO subject_grade_components (
			subject_id,
			name,
			percentage,
			is_active,
			created_by
		) VALUES (
			:subject_id,
			:name,
			:percentage,
			:is_active,
			:created_by
		) ON CONFLICT (subject_id, name) DO UPDATE SET
			percentage = EXCLUDED.percentage,
			is_active = EXCLUDED.is_active,
			updated_by = EXCLUDED.created_by
	`

	deleteAllBySubjectIdExcludingNamesQuery = `
		DELETE FROM subject_grade_components WHERE subject_id = $1 AND name NOT IN (SELECT UNNEST($2::text[]))
	`

	deleteAllBySubjectIdQuery = `
		DELETE FROM subject_grade_components WHERE subject_id = $1
	`
)
