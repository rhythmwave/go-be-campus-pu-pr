package subject

const (
	getListQuery = `
		SELECT 
			s.id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			c.id AS curriculum_id,
			c.name AS curriculum_name,
			s.code,
			s.name,
			s.is_mandatory,
			s.semester_package,
			s.theory_credit,
			s.practicum_credit,
			s.field_practicum_credit,
			sc.id AS subject_category_id,
			sc.name AS subject_category_name,
			s.is_thesis,
			s.is_mbkm,
			s.total_lesson_plan
			{{PREREQUISITE_SELECT}}
			{{EQUIVALENT_SELECT}}
		FROM subjects s
		JOIN curriculums c ON c.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN subject_categories sc ON sc.id = s.subject_category_id
		{{PREREQUISITE_JOIN}}
		{{EQUIVALENT_JOIN}}
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM subjects s
		JOIN curriculums c ON c.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN subject_categories sc ON sc.id = s.subject_category_id
		{{PREREQUISITE_JOIN}}
		{{EQUIVALENT_JOIN}}
	`

	getDetailQuery = `
		SELECT 
			s.id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			c.id AS curriculum_id,
			c.name AS curriculum_name,
			s.code,
			s.name,
			s.short_name,
			s.english_name,
			s.english_short_name,
			s.is_mandatory,
			s.trait,
			s.type,
			sc.id AS subject_category_id,
			sc.name AS subject_category_name,
			s.curriculum_type,
			s.theory_credit,
			s.practicum_credit,
			s.field_practicum_credit,
			s.semester_package,
			s.repeat_course_limit,
			s.is_active,
			s.has_lecture_unit,
			s.has_teaching_material,
			s.has_lecture_summary,
			sl.id AS supporting_lecturer_id,
			sl.name AS supporting_lecturer_name,
			s.start_date,
			s.end_date,
			s.minimum_passing_grade_point,
			s.minimum_mandatory_credit_taken,
			s.minimum_optional_credit_taken,
			s.minimum_total_credit_taken,
			s.minimum_mandatory_credit_passed,
			s.minimum_optional_credit_passed,
			s.minimum_total_credit_passed,
			s.minimum_gpa,
			s.abstraction,
			s.syllabus_path,
			s.syllabus_path_type,
			s.is_thesis,
			s.is_mbkm
		FROM subjects s
		JOIN curriculums c ON c.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN subject_categories sc ON sc.id = s.subject_category_id
		LEFT JOIN lecturers sl ON sl.id = s.supporting_lecturer_id
		WHERE s.id = $1
	`

	getDetailByIdsQuery = `
		SELECT 
			s.id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			c.id AS curriculum_id,
			c.name AS curriculum_name,
			s.code,
			s.name,
			s.short_name,
			s.english_name,
			s.english_short_name,
			s.is_mandatory,
			s.trait,
			s.type,
			sc.id AS subject_category_id,
			sc.name AS subject_category_name,
			s.curriculum_type,
			s.theory_credit,
			s.practicum_credit,
			s.field_practicum_credit,
			s.semester_package,
			s.repeat_course_limit,
			s.is_active,
			s.has_lecture_unit,
			s.has_teaching_material,
			s.has_lecture_summary,
			sl.id AS supporting_lecturer_id,
			sl.name AS supporting_lecturer_name,
			s.start_date,
			s.end_date,
			s.minimum_passing_grade_point,
			s.minimum_mandatory_credit_taken,
			s.minimum_optional_credit_taken,
			s.minimum_total_credit_taken,
			s.minimum_mandatory_credit_passed,
			s.minimum_optional_credit_passed,
			s.minimum_total_credit_passed,
			s.minimum_gpa,
			s.abstraction,
			s.syllabus_path,
			s.syllabus_path_type,
			s.is_thesis,
			s.is_mbkm
		FROM subjects s
		JOIN curriculums c ON c.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN subject_categories sc ON sc.id = s.subject_category_id
		LEFT JOIN lecturers sl ON sl.id = s.supporting_lecturer_id
		WHERE s.id IN (SELECT UNNEST($1::uuid[]))
	`

	getThesisByCurriculumIdQuery = `
		SELECT 
			s.id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			c.id AS curriculum_id,
			c.name AS curriculum_name,
			s.code,
			s.name,
			s.short_name,
			s.english_name,
			s.english_short_name,
			s.is_mandatory,
			s.trait,
			s.type,
			sc.id AS subject_category_id,
			sc.name AS subject_category_name,
			s.curriculum_type,
			s.theory_credit,
			s.practicum_credit,
			s.field_practicum_credit,
			s.semester_package,
			s.repeat_course_limit,
			s.is_active,
			s.has_lecture_unit,
			s.has_teaching_material,
			s.has_lecture_summary,
			sl.id AS supporting_lecturer_id,
			sl.name AS supporting_lecturer_name,
			s.start_date,
			s.end_date,
			s.minimum_passing_grade_point,
			s.minimum_mandatory_credit_taken,
			s.minimum_optional_credit_taken,
			s.minimum_total_credit_taken,
			s.minimum_mandatory_credit_passed,
			s.minimum_optional_credit_passed,
			s.minimum_total_credit_passed,
			s.minimum_gpa,
			s.abstraction,
			s.syllabus_path,
			s.syllabus_path_type,
			s.is_thesis,
			s.is_mbkm
		FROM subjects s
		JOIN curriculums c ON c.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = c.study_program_id
		JOIN subject_categories sc ON sc.id = s.subject_category_id
		LEFT JOIN lecturers sl ON sl.id = s.supporting_lecturer_id
		WHERE s.curriculum_id = $1 AND s.is_thesis IS true
	`

	createQuery = `
		INSERT INTO subjects (
			curriculum_id,
			code,
			name,
			short_name,
			english_name,
			english_short_name,
			is_mandatory,
			trait,
			type,
			subject_category_id,
			curriculum_type,
			theory_credit,
			practicum_credit,
			field_practicum_credit,
			semester_package,
			repeat_course_limit,
			is_active,
			has_lecture_unit,
			has_teaching_material,
			has_lecture_summary,
			supporting_lecturer_id,
			start_date,
			end_date,
			minimum_passing_grade_point,
			minimum_mandatory_credit_taken,
			minimum_optional_credit_taken,
			minimum_total_credit_taken,
			minimum_mandatory_credit_passed,
			minimum_optional_credit_passed,
			minimum_total_credit_passed,
			minimum_gpa,
			abstraction,
			syllabus_path,
			syllabus_path_type,
			is_thesis,
			is_mbkm,
			created_by
		) VALUES (
			:curriculum_id,
			:code,
			:name,
			:short_name,
			:english_name,
			:english_short_name,
			:is_mandatory,
			:trait,
			:type,
			:subject_category_id,
			:curriculum_type,
			:theory_credit,
			:practicum_credit,
			:field_practicum_credit,
			:semester_package,
			:repeat_course_limit,
			:is_active,
			:has_lecture_unit,
			:has_teaching_material,
			:has_lecture_summary,
			:supporting_lecturer_id,
			:start_date,
			:end_date,
			:minimum_passing_grade_point,
			:minimum_mandatory_credit_taken,
			:minimum_optional_credit_taken,
			:minimum_total_credit_taken,
			:minimum_mandatory_credit_passed,
			:minimum_optional_credit_passed,
			:minimum_total_credit_passed,
			:minimum_gpa,
			:abstraction,
			:syllabus_path,
			:syllabus_path_type,
			:is_thesis,
			:is_mbkm,
			:created_by
		)
	`

	updateQuery = `
		UPDATE subjects SET
			curriculum_id = :curriculum_id,
			code = :code,
			name = :name,
			short_name = :short_name,
			english_name = :english_name,
			english_short_name = :english_short_name,
			is_mandatory = :is_mandatory,
			trait = :trait,
			type = :type,
			subject_category_id = :subject_category_id,
			curriculum_type = :curriculum_type,
			theory_credit = :theory_credit,
			practicum_credit = :practicum_credit,
			field_practicum_credit = :field_practicum_credit,
			semester_package = :semester_package,
			repeat_course_limit = :repeat_course_limit,
			is_active = :is_active,
			has_lecture_unit = :has_lecture_unit,
			has_teaching_material = :has_teaching_material,
			has_lecture_summary = :has_lecture_summary,
			supporting_lecturer_id = :supporting_lecturer_id,
			start_date = :start_date,
			end_date = :end_date,
			minimum_passing_grade_point = :minimum_passing_grade_point,
			minimum_mandatory_credit_taken = :minimum_mandatory_credit_taken,
			minimum_optional_credit_taken = :minimum_optional_credit_taken,
			minimum_total_credit_taken = :minimum_total_credit_taken,
			minimum_mandatory_credit_passed = :minimum_mandatory_credit_passed,
			minimum_optional_credit_passed = :minimum_optional_credit_passed,
			minimum_total_credit_passed = :minimum_total_credit_passed,
			minimum_gpa = :minimum_gpa,
			abstraction = :abstraction,
			syllabus_path = :syllabus_path,
			syllabus_path_type = :syllabus_path_type,
			is_mbkm = :is_mbkm,
			is_thesis = :is_thesis
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM subjects WHERE id = $1
	`
)
