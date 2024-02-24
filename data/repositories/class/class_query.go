package class

const (
	getListQuery = `
		SELECT 
			c.id,
			c.name,
			cu.study_program_id,
			cu.id AS curriculum_id,
			cu.name AS curriculum_name,
			c.semester_id,
			se.semester_start_year,
			se.semester_type,
			s.id AS subject_id,
			s.code AS subject_code,
			s.name AS subject_name,
			s.is_mandatory AS subject_is_mandatory,
			s.semester_package AS subject_semester_package,
			s.total_lesson_plan AS subject_total_lesson_plan,
			c.total_participant,
			c.maximum_participant,
			s.theory_credit AS subject_theory_credit,
			s.practicum_credit AS subject_practicum_credit,
			s.field_practicum_credit AS subject_field_practicum_credit,
			c.unapproved_study_plan,
			c.total_material,
			c.total_work,
			c.total_discussion,
			c.total_exam,
			c.total_event,
			c.total_lecture_plan,
			c.total_lecture_done,
			c.is_active,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.study_level_id,
			c.total_graded_participant,
			c.application_deadline
		FROM classes c
		JOIN semesters se ON se.id = c.semester_id
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM classes c
		JOIN semesters se ON se.id = c.semester_id
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
	`

	getBySubjectIdsSemesterIdQuery = `
		SELECT 
			c.id,
			c.name,
			cu.study_program_id,
			cu.id AS curriculum_id,
			cu.name AS curriculum_name,
			c.semester_id,
			s.id AS subject_id,
			s.code AS subject_code,
			s.name AS subject_name,
			s.is_mandatory AS subject_is_mandatory,
			s.semester_package AS subject_semester_package,
			s.total_lesson_plan AS subject_total_lesson_plan,
			c.total_participant,
			c.maximum_participant,
			s.theory_credit AS subject_theory_credit,
			s.practicum_credit AS subject_practicum_credit,
			s.field_practicum_credit AS subject_field_practicum_credit,
			c.unapproved_study_plan,
			c.total_material,
			c.total_work,
			c.total_discussion,
			c.total_exam,
			c.total_lecture_plan,
			c.total_lecture_done,
			c.is_active,
			dsp.study_level_id,
			c.total_graded_participant,
			c.application_deadline
		FROM classes c
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		WHERE s.id IN (SELECT UNNEST($1::uuid[])) AND c.semester_id = $2
		ORDER BY s.id, c.name
	`

	getDetailQuery = `
		SELECT 
			c.id,
			c.name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.type AS dikti_study_program_type,
			sl.id AS study_level_id,
			sl.short_name AS study_level_short_name,
			cu.id AS curriculum_id,
			cu.name AS curriculum_name,
			cu.year AS curriculum_year,
			se.id AS semester_id,
			se.semester_start_year,
			se.semester_type,
			se.is_active AS semester_is_active,
			se.grading_start_date,
			se.grading_end_date,
			s.id AS subject_id,
			s.code AS subject_code,
			s.name AS subject_name,
			s.theory_credit AS subject_theory_credit,
			s.practicum_credit AS subject_practicum_credit,
			s.field_practicum_credit AS subject_field_practicum_credit,
			c.scope,
			c.is_online,
			c.is_offline,
			c.minimum_participant,
			c.maximum_participant,
			c.total_participant,
			c.remarks,
			c.is_active,
			s.semester_package,
			s.theory_credit,
			s.practicum_credit,
			s.field_practicum_credit,
			s.minimum_passing_grade_point AS subject_minimum_passing_grade_point,
			s.is_mandatory AS subject_is_mandatory,
			c.total_material,
			c.total_work,
			c.total_event,
			c.total_discussion,
			c.total_exam,
			c.total_lecture_plan,
			c.total_lecture_done,
			c.application_deadline,
			s.is_mbkm AS subject_is_mbkm
			%s
		FROM classes c
		JOIN semesters se ON se.id = c.semester_id
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		%s
		WHERE c.id = $1
	`

	getDetailByIdsQuery = `
		SELECT 
			c.id,
			c.name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.type AS dikti_study_program_type,
			sl.id AS study_level_id,
			sl.short_name AS study_level_short_name,
			cu.id AS curriculum_id,
			cu.name AS curriculum_name,
			cu.year AS curriculum_year,
			se.id AS semester_id,
			se.semester_start_year,
			se.semester_type,
			se.is_active AS semester_is_active,
			se.grading_start_date,
			se.grading_end_date,
			s.id AS subject_id,
			s.code AS subject_code,
			s.name AS subject_name,
			s.theory_credit AS subject_theory_credit,
			s.practicum_credit AS subject_practicum_credit,
			s.field_practicum_credit AS subject_field_practicum_credit,
			c.scope,
			c.is_online,
			c.is_offline,
			c.minimum_participant,
			c.maximum_participant,
			c.total_participant,
			c.remarks,
			c.is_active,
			s.semester_package,
			s.theory_credit,
			s.practicum_credit,
			s.field_practicum_credit,
			s.minimum_passing_grade_point AS subject_minimum_passing_grade_point,
			s.is_mandatory AS subject_is_mandatory,
			c.total_material,
			c.total_work,
			c.total_event,
			c.total_discussion,
			c.total_exam,
			c.total_lecture_plan,
			c.total_lecture_done,
			c.application_deadline,
			s.is_mbkm AS subject_is_mbkm
			%s
		FROM classes c
		JOIN semesters se ON se.id = c.semester_id
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		%s
		WHERE c.id IN (SELEcT (UNNEST($1::uuid[])))
	`

	createQuery = `
		INSERT INTO classes (
			subject_id,
			semester_id,
			name,
			scope,
			is_online,
			is_offline,
			minimum_participant,
			maximum_participant,
			remarks,
			application_deadline,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id
	`

	updateQuery = `
		UPDATE classes SET
			subject_id = $1,
			name = $2,
			scope = $3,
			is_online = $4,
			is_offline = $5,
			minimum_participant = $6,
			maximum_participant = $7,
			remarks = $8,
			application_deadline = $9,
			updated_by = $10
		WHERE id = $11
	`

	upsertMaximumParticipantQuery = `
		INSERT INTO classes (
			id,
			subject_id,
			semester_id,
			name,
			scope,
			is_online,
			is_offline,
			minimum_participant,
			maximum_participant,
			remarks,
			created_by
		) VALUES (
			:id,
			:subject_id,
			:semester_id,
			:name,
			:scope,
			:is_online,
			:is_offline,
			:minimum_participant,
			:maximum_participant,
			:remarks,
			:created_by
		) ON CONFLICT (id) DO UPDATE SET
			maximum_participant = EXCLUDED.maximum_participant,
			updated_by = EXCLUDED.created_by
	`

	updateActivationQuery = `
		UPDATE classes SET is_active = $1 WHERE id = $2
	`

	deleteQuery = `
		DELETE FROM classes WHERE id = $1
	`

	getClassLecturerByClassIdsQuery = `
		SELECT
			cl.id,
			cl.class_id,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cl.is_grading_responsible
		FROM class_lecturers cl
		JOIN lecturers l ON l.id = cl.lecturer_id
		WHERE cl.class_id IN (SELECT UNNEST($1::uuid[]))
	`

	getClassLecturersBySemesterIdLecturerIdQuery = `
		SELECT
			cl.id,
			cl.class_id,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			cl.is_grading_responsible
		FROM class_lecturers cl
		JOIN lecturers l ON l.id = cl.lecturer_id
		JOIN classes c ON c.id = cl.class_id
		WHERE c.semester_id = $1 AND l.id = $2
	`

	deleteClassLecturerExcludingLecturerIdsQuery = `
		DELETE FROM class_lecturers WHERE class_id = $1 %s
	`

	upsertClassLecturerQuery = `
		INSERT INTO class_lecturers (
			class_id,
			lecturer_id,
			is_grading_responsible,
			created_by
		) VALUES (
			:class_id,
			:lecturer_id,
			:is_grading_responsible,
			:created_by
		) ON CONFLICT (class_id, lecturer_id) DO UPDATE SET
			is_grading_responsible = EXCLUDED.is_grading_responsible,
			updated_by = EXCLUDED.created_by
	`

	duplicateQuery = `
		INSERT INTO classes (
			subject_id,
			semester_id,
			name,
			scope,
			is_online,
			is_offline,
			minimum_participant,
			maximum_participant,
			remarks,
			created_by
		) SELECT
			c.subject_id,
			$2,
			c.name,
			c.scope,
			c.is_online,
			c.is_offline,
			c.minimum_participant,
			c.maximum_participant,
			c.remarks,
			$3
		FROM classes c
		WHERE c.semester_id = $1
	`

	duplicateLecturerQuery = `
		INSERT INTO class_lecturers (
			class_id,
			lecturer_id,
			is_grading_responsible,
			created_by
		) SELECT
			ct.id,
			cl.lecturer_id,
			cl.is_grading_responsible,
			$3
		FROM class_lecturers cl
		JOIN classes cf ON cf.id = class_id AND cf.semester_id = $1
		JOIN classes ct ON ct.subject_id = cf.subject_id AND ct.name = cf.name AND ct.semester_id = $2
	`

	getActiveBySemesterIdQuery = `
		SELECT 
			c.id,
			c.name,
			cu.study_program_id,
			cu.id AS curriculum_id,
			cu.name AS curriculum_name,
			c.semester_id,
			s.id AS subject_id,
			s.code AS subject_code,
			s.name AS subject_name,
			s.is_mandatory AS subject_is_mandatory,
			s.semester_package AS subject_semester_package,
			s.total_lesson_plan AS subject_total_lesson_plan,
			c.total_participant,
			c.maximum_participant,
			s.theory_credit AS subject_theory_credit,
			s.practicum_credit AS subject_practicum_credit,
			s.field_practicum_credit AS subject_field_practicum_credit,
			c.total_material,
			c.total_work,
			c.total_discussion,
			c.total_exam,
			c.total_lecture_plan,
			c.total_lecture_done,
			c.is_active,
			dsp.study_level_id,
			c.total_graded_participant,
			c.application_deadline
		FROM classes c
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		WHERE c.is_active IS true AND c.semester_id = $1
	`

	inactivateClassesQuery = `
		UPDATE classes SET is_active = false WHERE id IN (SELECT UNNEST($1::uuid[]))
	`

	getThesisClassQuery = `
		SELECT 
			c.id,
			c.name,
			cu.study_program_id,
			cu.id AS curriculum_id,
			cu.name AS curriculum_name,
			c.semester_id,
			s.id AS subject_id,
			s.code AS subject_code,
			s.name AS subject_name,
			s.is_mandatory AS subject_is_mandatory,
			s.semester_package AS subject_semester_package,
			s.total_lesson_plan AS subject_total_lesson_plan,
			c.total_participant,
			c.maximum_participant,
			s.theory_credit AS subject_theory_credit,
			s.practicum_credit AS subject_practicum_credit,
			s.field_practicum_credit AS subject_field_practicum_credit,
			c.unapproved_study_plan,
			c.total_material,
			c.total_work,
			c.total_discussion,
			c.total_exam,
			c.total_lecture_plan,
			c.total_lecture_done,
			c.is_active,
			dsp.study_level_id,
			c.total_graded_participant,
			c.application_deadline
		FROM classes c
		JOIN student_classes sc ON sc.class_id = c.id AND sc.student_id = $1
		JOIN study_plans spl ON spl.id = sc.study_plan_id
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		WHERE c.semester_id = $2 AND s.is_thesis IS true
	`

	upsertThesisClassQuery = `
		INSERT INTO classes (
			subject_id,
			semester_id,
			name,
			created_by
		) SELECT
			$1,
			$3,
			l.id_national_lecturer,
			$4
		FROM lecturers l
		WHERE l.id = $2
		ON CONFLICT (subject_id, semester_id, name) DO UPDATE SET id = EXCLUDED.id
		RETURNING id
	`

	upsertThesisClassLecturer = `
		INSERT INTO class_lecturers (
			class_id,
			lecturer_id,
			is_grading_responsible,
			created_by
		) VALUES (
			$1,
			$2,
			true,
			$3
		) ON CONFLICT (class_id, lecturer_id) DO NOTHING
	`
)
