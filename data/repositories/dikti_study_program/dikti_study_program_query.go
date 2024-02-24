package dikti_study_program

const (
	getListQuery = `
		SELECT 
			dsp.id, 
			dsp.code,
			dsp.name,
			sl.short_name AS study_level_short_name,
			sl.name AS study_level_name
		FROM dikti_study_programs dsp
		JOIN study_levels sl ON sl.id = dsp.study_level_id
	`

	countListQuery = `
		SELECT COUNT(1) 
		FROM dikti_study_programs dsp
		JOIN study_levels sl ON sl.id = dsp.study_level_id
	`
)
