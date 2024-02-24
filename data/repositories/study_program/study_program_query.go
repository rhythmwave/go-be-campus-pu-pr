package study_program

const (
	getListQuery = `
		SELECT
			sp.id,
			sp.name,
			sl.short_name AS study_level_short_name,
			sl.name AS study_level_name,
			dsp.type AS dikti_study_program_type,
			dsp.code AS dikti_study_program_code,
			ac.accreditation,
			c.year AS active_curriculum_year,
			sp.degree,
			sp.short_degree,
			sp.english_degree
		FROM study_programs sp
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		LEFT JOIN accreditations ac ON ac.study_program_id = sp.id AND ac.is_active IS true
		LEFT JOIN curriculums c ON c.study_program_id = sp.id AND c.is_active IS true
		%s
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM study_programs sp
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		%s
	`

	getDetailQuery = `
		SELECT 
			sp.id,
			dsp.id AS dikti_study_program_id,
			dsp.code AS dikti_study_program_code,
			dsp.name AS dikti_study_program_name,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			sl.name AS study_level_name,
			sp.name,
			sp.english_name,
			sp.short_name,
			sp.english_short_name,
			sp.administrative_unit,
			f.id AS faculty_id,
			f.name AS faculty_name,
			m.id AS major_id,
			m.name AS major_name,
			sp.address,
			sp.phone_number,
			sp.fax,
			sp.email,
			sp.website,
			sp.contact_person,
			sp.curiculum_review_frequency,
			sp.curiculum_review_method,
			sp.establishment_date,
			sp.is_active,
			sp.start_semester,
			sp.operational_permit_number,
			sp.operational_permit_date,
			sp.operational_permit_due_date,
			l.id AS head_lecturer_id,
			l.name AS head_lecturer_name,
			l.mobile_phone_number AS head_lecturer_mobile_phone_number,
			sp.operator_name,
			sp.operator_phone_number,
			sp.minimum_graduation_credit,
			sp.minimum_thesis_credit
		FROM study_programs sp
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN majors m ON m.id = sp.major_id
		JOIN faculties f ON f.id = m.faculty_id
		LEFT JOIN lecturers l ON l.id = sp.head_lecturer_id
		%s
		WHERE sp.id = $1
	`

	createQuery = `
		INSERT INTO study_programs (
			dikti_study_program_id,
			name,
			english_name,
			short_name,
			english_short_name,
			administrative_unit,
			major_id,
			address,
			phone_number,
			fax,
			email,
			website,
			contact_person,
			curiculum_review_frequency,
			curiculum_review_method,
			establishment_date,
			is_active,
			start_semester,
			operational_permit_number,
			operational_permit_date,
			operational_permit_due_date,
			head_lecturer_id,
			operator_name,
			operator_phone_number,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25);
	`

	updateQuery = `
		UPDATE study_programs SET
			dikti_study_program_id = $1,
			name = $2,
			english_name = $3,
			short_name = $4,
			english_short_name = $5,
			administrative_unit = $6,
			major_id = $7,
			address = $8,
			phone_number = $9,
			fax = $10,
			email = $11,
			website = $12,
			contact_person = $13,
			curiculum_review_frequency = $14,
			curiculum_review_method = $15,
			establishment_date = $16,
			is_active = $17,
			start_semester = $18,
			operational_permit_number = $19,
			operational_permit_date = $20,
			operational_permit_due_date = $21,
			head_lecturer_id = $22,
			operator_name = $23,
			operator_phone_number = $24,
			updated_by = $25,
			minimum_graduation_credit = COALESCE($26, minimum_graduation_credit),
			minimum_thesis_credit = COALESCE($27, minimum_thesis_credit)
		WHERE id = $28
	`

	updateDegreeQuery = `
		UPDATE study_programs SET
			degree = $1,
			short_degree = $2,
			english_degree = $3,
			updated_by = $4
		WHERE id = $5
	`

	deleteQuery = `
		DELETE FROM study_programs WHERE id = $1
	`

	getByRoleIdsQuery = `
		SELECT
			sp.id,
			sp.name,
			sl.short_name AS study_level_short_name,
			dsp.type AS dikti_study_program_type,
			rsp.role_id
		FROM study_programs sp
		JOIN role_study_program rsp ON rsp.study_program_id = sp.id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		WHERE rsp.role_id IN (SELECT UNNEST($1::uuid[]))
		ORDER BY sp.name, rsp.role_id
	`
)
