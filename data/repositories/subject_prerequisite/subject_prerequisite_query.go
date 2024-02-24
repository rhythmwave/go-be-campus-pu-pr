package subject_prerequisite

const (
	getBySubjectIdQuery = `
		SELECT
			sp.id,
			sp.subject_id,
			sp.prerequisite_subject_id,
			ssp.code AS prerequisite_subject_code,
			ssp.name AS prerequisite_subject_name,
			sp.prerequisite_type,
			sp.minimum_grade_point
		FROM subject_prerequisites sp
		JOIN subjects ssp ON ssp.id = sp.prerequisite_subject_id
		WHERE sp.subject_id = $1
	`

	upsertQuery = `
		INSERT INTO subject_prerequisites (
			subject_id,
			prerequisite_subject_id,
			prerequisite_type,
			minimum_grade_point,
			created_by
		) VALUES (
			:subject_id,
			:prerequisite_subject_id,
			:prerequisite_type,
			:minimum_grade_point,
			:created_by
		) ON CONFLICT (subject_id, prerequisite_subject_id) DO UPDATE SET
			prerequisite_type = EXCLUDED.prerequisite_type,
			minimum_grade_point = EXCLUDED.minimum_grade_point,
			updated_by = EXCLUDED.created_by
	`

	deleteAllBySubjectIdExcludingPrerequisiteSubjectIdQuery = `
		DELETE FROM subject_prerequisites WHERE subject_id = $1 AND prerequisite_subject_id NOT IN (SELECT UNNEST($2::uuid[]))
	`

	deleteAllBySubjectIdQuery = `
		DELETE FROM subject_prerequisites WHERE subject_id = $1
	`
)
