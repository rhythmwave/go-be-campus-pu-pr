package subject_equivalence

const (
	upsertQuery = `
		INSERT INTO subject_equivalences (
			subject_id,
			equivalent_subject_id,
			created_by
		) VALUES ($1, $2, $3) 
		ON CONFLICT (subject_id, equivalent_curriculum_id)
		DO UPDATE SET 
			equivalent_subject_id = EXCLUDED.equivalent_subject_id,
			updated_by = EXCLUDED.created_by
	`

	deleteBySubjectIdAndEquivalentSubjectIdQuery = `
		DELETE FROM subject_equivalences WHERE subject_id = $1 AND equivalent_subject_id = $2
	`
)
