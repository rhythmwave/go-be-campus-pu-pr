package student_class

const (
	getListQuery = `
		SELECT
			sc.id,
			c.id AS class_id,
			c.name AS class_name,
			s.id AS subject_id,
			s.code AS subject_code,
			s.name AS subject_name,
			se.id AS semester_id,
			se.semester_start_year,
			se.semester_type,
			s.theory_credit AS subject_theory_credit,
			s.practicum_credit AS subject_practicum_credit,
			s.field_practicum_credit AS subject_field_practicum_credit,
			sc.subject_repetition,
			sc.study_plan_id,
			sc.student_id,
			sc.subject_is_mandatory,
			sc.curriculum_id,
			sc.student_curriculum_id,
			sc.total_credit,
			sc.total_attendance,
			sc.total_sick,
			sc.total_leave,
			sc.total_awol,
			sc.grade_point,
			sc.grade_code,
			sc.mbkm_used_credit,
			ga.id AS graded_by_admin_id,
			ga.name AS graded_by_admin_name,
			gl.id AS graded_by_lecturer_id,
			gl.name AS graded_by_lecturer_name,
			sc.graded_at,
			c.total_lecture_done,
			l.id AS active_lecture_id,
			lp.is_attend AS active_lecture_has_attend,
			s.is_mbkm AS subject_is_mbkm
		FROM student_classes sc
		JOIN study_plans sp ON sp.id = sc.study_plan_id
		JOIN semesters se ON se.id = sp.semester_id
		JOIN classes c ON c.id = sc.class_id
		JOIN subjects s ON s.id = sc.subject_id
		LEFT JOIN admins ga ON ga.id = sc.graded_by_admin_id
		LEFT JOIN lecturers gl ON gl.id = sc.graded_by_lecturer_id
		LEFT JOIN lectures l ON l.class_id = c.id AND now() BETWEEN l.autonomous_participation_start_time	AND l.autonomous_participation_end_time
		LEFT JOIN lecture_participants lp ON lp.lecture_id = l.id AND lp.student_id = sc.student_id 
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM student_classes sc
		JOIN study_plans sp ON sp.id = sc.study_plan_id
		JOIN classes c ON c.id = sc.class_id
		JOIN subjects s ON s.id = sc.subject_id
		LEFT JOIN admins ga ON ga.id = sc.graded_by_admin_id
		LEFT JOIN lecturers gl ON gl.id = sc.graded_by_lecturer_id
	`

	getClassParticipantQuery = `
		SELECT 
			s.id AS student_id,
			s.nim_number AS student_nim_number,
			s.name AS student_name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.code AS dikti_study_program_code,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			sc.id AS student_class_id,
			sc.study_plan_id,
			sc.subject_id,
			su.name AS subject_name,
			sc.curriculum_id,
			sc.student_curriculum_id,
			sc.class_id,
			sc.total_attendance,
			sc.total_sick,
			sc.total_leave,
			sc.total_awol,
			sc.grade_point,
			sc.grade_code,
			ga.id AS graded_by_admin_id,
			ga.name AS graded_by_admin_name,
			gl.id AS graded_by_lecturer_id,
			gl.name AS graded_by_lecturer_name,
			sc.graded_at,
			sc.subject_repetition
			%s
		FROM student_classes sc
		JOIN students s ON s.id = sc.student_id
		JOIN subjects su ON su.id = sc.subject_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		LEFT JOIN admins ga ON ga.id = sc.graded_by_admin_id
		LEFT JOIN lecturers gl ON gl.id = sc.graded_by_lecturer_id
		%s
	`

	countClassParticipantQuery = `
		SELECT COUNT(1)
		FROM student_classes sc
		JOIN students s ON s.id = sc.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		%s
	`

	bulkCreateQuery = `
		INSERT INTO student_classes (
			study_plan_id,
			curriculum_id,
			student_curriculum_id,
			class_id
		) VALUES (
			:study_plan_id,
			:curriculum_id,
			:student_curriculum_id,
			:class_id
		) ON CONFLICT (student_id, class_id) DO NOTHING
	`

	bulkUpdateClassQuery = `
		INSERT INTO student_classes (
			id,
			study_plan_id,
			curriculum_id,
			student_curriculum_id,
			class_id
		) VALUES (
			:id,
			:study_plan_id,
			:curriculum_id,
			:student_curriculum_id,
			:class_id
		) ON CONFLICT (id) DO UPDATE SET
			class_id = EXCLUDED.class_id
	`

	bulkGradeStudentClassQuery = `
		INSERT INTO student_class_grades (
			class_id,
			student_id,
			class_grade_component_id,
			initial_grade,
			graded_by_admin_id,
			graded_by_lecturer_id
		) VALUES (
			:class_id,
			:student_id,
			:class_grade_component_id,
			:initial_grade,
			:graded_by_admin_id,
			:graded_by_lecturer_id
		) ON CONFLICT (student_id, class_id, class_grade_component_id) DO UPDATE SET
			initial_grade =  EXCLUDED.initial_grade,
			graded_by_admin_id =  EXCLUDED.graded_by_admin_id,
			graded_by_lecturer_id =  EXCLUDED.graded_by_lecturer_id
	`

	getStudentClassGradeByClassIdAndStudentIdsQuery = `
		SELECT
			scg.class_id,
			scg.student_id,
			cgc.id AS class_grade_component_id,
			cgc.name AS class_grade_component_name,
			scg.initial_grade,
			scg.final_grade
		FROM student_class_grades scg
		JOIN class_grade_components cgc ON cgc.id = scg.class_grade_component_id
		WHERE scg.class_id = $1 AND scg.student_id IN (SELECT UNNEST($2::uuid[]))
	`

	bulkDeleteExcludingClassIdsQuery = `
		DELETE FROM student_classes
		WHERE %s
	`

	updateMbkmConvertedCreditQuery = `
		UPDATE student_classes
		SET mbkm_used_credit = $1
		WHERE id = $2
	`
)
