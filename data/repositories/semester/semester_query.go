package semester

const (
	getListQuery = `
		SELECT 
			s.id,
			s.semester_start_year,
			s.semester_type,
			s.is_active,
			s.start_date,
			s.end_date,
			s.midterm_start_date,
			s.midterm_end_date,
			s.endterm_start_date,
			s.endterm_end_date
		FROM semesters s
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM semesters s
	`

	getByIdQuery = `
		SELECT 
			s.id,
			s.semester_start_year,
			s.semester_type,
			s.is_active,
			s.start_date,
			s.end_date,
			s.midterm_start_date,
			s.midterm_end_date,
			s.endterm_start_date,
			s.endterm_end_date,
			s.study_plan_input_start_date,
			s.study_plan_input_end_date,
			s.study_plan_approval_start_date,
			s.study_plan_approval_end_date,
			s.grading_start_date,
			s.grading_end_date,
			rs.id AS reference_semester_id,
			rs.semester_start_year AS reference_semester_start_year,
			rs.semester_type AS reference_semester_type,
			s.check_minimum_gpa,
			s.check_passed_credit,
			s.default_credit
		FROM semesters s
		LEFT JOIN semesters rs ON rs.id = s.reference_semester_id
		WHERE s.id = $1
	`

	getActiveQuery = `
		SELECT 
			s.id,
			s.semester_start_year,
			s.semester_type,
			s.is_active,
			s.start_date,
			s.end_date,
			s.midterm_start_date,
			s.midterm_end_date,
			s.endterm_start_date,
			s.endterm_end_date,
			s.study_plan_input_start_date,
			s.study_plan_input_end_date,
			s.study_plan_approval_start_date,
			s.study_plan_approval_end_date,
			s.grading_start_date,
			s.grading_end_date,
			rs.id AS reference_semester_id,
			rs.semester_start_year AS reference_semester_start_year,
			rs.semester_type AS reference_semester_type,
			s.check_minimum_gpa,
			s.check_passed_credit,
			s.default_credit
		FROM semesters s
		LEFT JOIN semesters rs ON rs.id = s.reference_semester_id
		WHERE s.is_active IS true
	`

	getCurriculumBySemesterIdsQuery = `
		SELECT
			sc.semester_id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			c.id AS curriculum_id,
			c.name AS curriculum_name
		FROM semester_curriculum sc
		JOIN curriculums c ON c.id = sc.curriculum_id
		JOIN study_programs sp ON sp.id = c.study_program_id
		WHERE sc.semester_id IN (SELECT UNNEST($1::uuid[]))
	`

	createQuery = `
		INSERT INTO semesters (
			semester_start_year,
			semester_type,
			start_date,
			end_date,
			midterm_start_date,
			midterm_end_date,
			endterm_start_date,
			endterm_end_date,
			study_plan_input_start_date,
			study_plan_input_end_date,
			study_plan_approval_start_date,
			study_plan_approval_end_date,
			reference_semester_id,
			check_minimum_gpa,
			check_passed_credit,
			default_credit,
			grading_start_date,
			grading_end_date,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
		RETURNING id
	`

	updateQuery = `
		UPDATE semesters SET
			semester_start_year = $1,
			semester_type = $2,
			start_date = $3,
			end_date = $4,
			study_plan_input_start_date = $5,
			study_plan_input_end_date = $6,
			study_plan_approval_start_date = $7,
			study_plan_approval_end_date = $8,
			reference_semester_id = $9,
			check_minimum_gpa = $10,
			check_passed_credit = $11,
			default_credit = $12,
			updated_by = $13,
			midterm_start_date = $14,
			midterm_end_date = $15,
			endterm_start_date = $16,
			endterm_end_date = $17,
			grading_start_date = $18,
			grading_end_date = $19
		WHERE id = $20
	`

	deleteCurriculumSemesterExcludingCurriculumIdQuery = `
		DELETE FROM semester_curriculum WHERE semester_id = $1
	`

	upsertCurriculumQuery = `
		INSERT INTO semester_curriculum (
			semester_id,
			curriculum_id,
			created_by
		) VALUES (
			:semester_id,
			:curriculum_id,
			:created_by
		) ON CONFLICT (semester_id, curriculum_id) DO NOTHING
	`

	updateActivationQuery = `
		UPDATE semesters SET is_active = $1 WHERE id = $2
	`

	deleteQuery = `
		DELETE FROM semesters WHERE id = $1
	`

	autoSetActiveQuery = `
		UPDATE semesters SET is_active = true WHERE start_date = DATE(now()) RETURNING id;
	`

	getPreviousSemesterQuery = `
		SELECT 
			s.id,
			s.semester_start_year,
			s.semester_type,
			s.is_active,
			s.start_date,
			s.end_date,
			s.midterm_start_date,
			s.midterm_end_date,
			s.endterm_start_date,
			s.endterm_end_date,
			s.study_plan_input_start_date,
			s.study_plan_input_end_date,
			s.study_plan_approval_start_date,
			s.study_plan_approval_end_date,
			s.grading_start_date,
			s.grading_end_date,			
			rs.id AS reference_semester_id,
			rs.semester_start_year AS reference_semester_start_year,
			rs.semester_type AS reference_semester_type,
			s.check_minimum_gpa,
			s.check_passed_credit,
			s.default_credit
		FROM semesters s
		JOIN semesters cs ON cs.id = $1
		LEFT JOIN semesters rs ON rs.id = s.reference_semester_id
		WHERE s.end_date < cs.start_date
		ORDER BY s.end_date DESC
		LIMIT 1
	`
)
