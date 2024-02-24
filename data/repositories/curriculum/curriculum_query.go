package curriculum

const (
	getListQuery = `
		SELECT 
		 c.id,
		 c.study_program_id,
		 dsp.code AS dikti_study_program_code,
		 sp.name AS study_program_name,
		 c.name,
		 c.year,
		 c.ideal_study_period,
		 c.maximum_study_period,
		 c.is_active,
		c.total_subject,
		c.total_subject_with_prerequisite,
		c.total_subject_with_equivalence
		FROM curriculums c
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM curriculums c
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
	`

	getDetailQuery = `
		SELECT 
			c.id,
			c.study_program_id,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			c.name,
			c.year,
			c.rector_decision_number,
			c.rector_decision_date,
			c.aggreeing_party,
			c.aggreement_date,
			c.ideal_study_period,
			c.maximum_study_period,
			c.remarks,
			c.is_active,
			c.final_score_determinant
		FROM curriculums c
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE c.id = $1
	`

	getActiveByStudyProgramIdQuery = `
		SELECT 
			c.id,
			c.study_program_id,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			c.name,
			c.year,
			c.rector_decision_number,
			c.rector_decision_date,
			c.aggreeing_party,
			c.aggreement_date,
			c.ideal_study_period,
			c.maximum_study_period,
			c.remarks,
			c.is_active,
			c.final_score_determinant
		FROM curriculums c
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE c.study_program_id = $1 AND c.is_active IS true
	`

	getActiveQuery = `
		SELECT
			c.id,
			c.study_program_id,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			c.name,
			c.year,
			c.rector_decision_number,
			c.rector_decision_date,
			c.aggreeing_party,
			c.aggreement_date,
			c.ideal_study_period,
			c.maximum_study_period,
			c.remarks,
			c.is_active,
			c.final_score_determinant
		FROM curriculums c
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE c.is_active IS true
	`

	createQuery = `
		INSERT INTO curriculums (
			study_program_id,
			name,
			year,
			rector_decision_number,
			rector_decision_date,
			aggreeing_party,
			aggreement_date,
			ideal_study_period,
			maximum_study_period,
			remarks,
			is_active,
			final_score_determinant,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	`

	updateQuery = `
		UPDATE curriculums SET
			name = $1,
			year = $2,
			rector_decision_number = $3,
			rector_decision_date = $4,
			aggreeing_party = $5,
			aggreement_date = $6,
			ideal_study_period = $7,
			maximum_study_period = $8,
			remarks = $9,
			is_active = $10,
			final_score_determinant = $11,
			updated_by = $12
		WHERE id = $13
	`

	deleteQuery = `
		DELETE FROM curriculums WHERE id = $1
	`
)
