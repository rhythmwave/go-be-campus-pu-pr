package lecturer_mutation

const (
	getListQuery = `
		SELECT
			lm.id,
			l.name,
			l.id_national_lecturer,
			l.front_title,
			l.back_degree,
			s.semester_start_year,
			s.semester_type,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			lm.mutation_date,
			lm.decision_number,
			lm.destination
		FROM lecturer_mutations lm
		JOIN lecturers l ON l.id = lm.lecturer_id
		JOIN semesters s ON s.id = lm.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM lecturer_mutations lm
		JOIN lecturers l ON l.id = lm.lecturer_id
		JOIN semesters s ON s.id = lm.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
	`

	getDetailQuery = `
		SELECT
			lm.id,
			l.name,
			l.id_national_lecturer,
			l.front_title,
			l.back_degree,
			s.semester_start_year,
			s.semester_type,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			lm.mutation_date,
			lm.decision_number,
			lm.destination
		FROM lecturer_mutations lm
		JOIN lecturers l ON l.id = lm.lecturer_id
		JOIN semesters s ON s.id = lm.semester_id
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE lm.id = $1
	`

	createQuery = `
		INSERT INTO lecturer_mutations (
			lecturer_id,
			semester_id,
			mutation_date,
			decision_number,
			destination,
			created_by
		) VALUES (
			:lecturer_id,
			:semester_id,
			:mutation_date,
			:decision_number,
			:destination,
			:created_by
		)
	`
)
