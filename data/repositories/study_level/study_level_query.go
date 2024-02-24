package study_level

const (
	getListQuery = `
		SELECT 
			sl.id, 
			sl.name,
			sl.short_name,
			sl.kkni_qualification,
			sl.acceptance_requirement,
			sl.further_education_level,
			sl.professional_status,
			sl.course_language
		FROM study_levels sl
	`

	countListQuery = `
		SELECT COUNT(1) 
		FROM study_levels sl
	`

	getDetailQuery = `
		SELECT 
			sl.id, 
			sl.name,
			sl.short_name,
			sl.kkni_qualification,
			sl.acceptance_requirement,
			sl.further_education_level,
			sl.professional_status,
			sl.course_language
		FROM study_levels sl
		WHERE sl.id = $1
	`

	updateSkpiQuery = `
		UPDATE study_levels SET
			kkni_qualification = $1,
			acceptance_requirement = $2,
			further_education_level = $3,
			professional_status = $4,
			course_language = $5,
			updated_by = $6
		WHERE id = $7
	`
)
