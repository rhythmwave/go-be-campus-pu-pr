package lecture

const (
	getListQuery = `
		SELECT
			l.id,
			l.lecture_plan_date,
			l.lecture_plan_day_of_week,
			l.lecture_plan_start_time,
			l.lecture_plan_end_time,
			l.lecture_actual_date,
			l.lecture_actual_day_of_week,
			l.lecture_actual_start_time,
			l.lecture_actual_end_time,
			lr.id AS lecturer_id,
			lr.name AS lecturer_name,
			l.foreign_lecturer_name,
			l.foreign_lecturer_source_instance,
			l.is_original_lecturer,
			l.is_manual_participation,
			l.autonomous_participation_start_time,
			l.autonomous_participation_end_time,
			l.attending_participant,
			c.id AS class_id,
			c.name AS class_name,
			r.id AS room_id,
			r.name AS room_name,
			b.id AS building_id,
			b.name AS building_name,
			l.is_midterm_exam,
			l.is_endterm_exam,
			l.is_theory_exam,
			l.is_practicum_exam,
			l.is_field_practicum_exam,
			s.code AS subject_code,
			s.name AS subject_name,
			l.total_participant,
			l.updated_at
		FROM lectures l
		LEFT JOIN lecturers lr ON lr.id = l.lecturer_id
		JOIN classes c ON c.id = l.class_id
		JOIN subjects s ON s.id = c.subject_id
		JOIN rooms r ON r.id = l.room_id
		JOIN buildings b ON b.id = r.building_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM lectures l
		LEFT JOIN lecturers lr ON lr.id = l.lecturer_id
	`

	getDetailQuery = `
		SELECT
			l.id,
			l.class_id,
			l.lecture_plan_date,
			l.lecture_plan_day_of_week,
			l.lecture_plan_start_time,
			l.lecture_plan_end_time,
			l.lecture_actual_date,
			l.lecture_actual_day_of_week,
			l.lecture_actual_start_time,
			l.lecture_actual_end_time,
			l.lecturer_id,
			l.foreign_lecturer_name,
			l.foreign_lecturer_source_instance,
			l.lecture_theme,
			l.lecture_subject,
			l.remarks,
			c.name AS class_name,
			s.id AS subject_id,
			s.name AS subject_name,
			se.id AS semester_id,
			se.semester_start_year,
			se.semester_type,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			r.id AS room_id,
			r.name AS room_name,
			l.is_exam,
			l.is_midterm_exam,
			l.is_endterm_exam,
			l.is_theory_exam,
			l.is_practicum_exam,
			l.is_field_practicum_exam,
			l.autonomous_participation_start_time,
			l.autonomous_participation_end_time
		FROM lectures l
		JOIN classes c ON c.id = l.class_id
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN semesters se ON se.id = c.semester_id
		JOIN rooms r ON r.id = l.room_id
		WHERE l.id = $1
	`

	getByClassIdsQuery = `
		SELECT
			l.id,
			l.class_id,
			l.lecture_plan_date,
			l.lecture_plan_day_of_week,
			l.lecture_plan_start_time,
			l.lecture_plan_end_time,
			l.lecture_actual_date,
			l.lecture_actual_day_of_week,
			l.lecture_actual_start_time,
			l.lecture_actual_end_time,
			l.lecturer_id,
			l.foreign_lecturer_name,
			l.foreign_lecturer_source_instance,
			l.lecture_theme,
			l.lecture_subject,
			l.remarks,
			c.name AS class_name,
			s.id AS subject_id,
			s.name AS subject_name,
			se.id AS semester_id,
			se.semester_start_year,
			se.semester_type,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			r.id AS room_id,
			r.name AS room_name,
			l.is_exam,
			l.is_midterm_exam,
			l.is_endterm_exam,
			l.is_theory_exam,
			l.is_practicum_exam,
			l.is_field_practicum_exam,
			l.autonomous_participation_start_time,
			l.autonomous_participation_end_time
		FROM lectures l
		JOIN classes c ON c.id = l.class_id
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN semesters se ON se.id = c.semester_id
		JOIN rooms r ON r.id = l.room_id
		WHERE l.class_id IN (SELECT UNNEST($1::uuid[]))
	`

	bulkCreateQuery = `
		INSERT INTO lectures (
			id,
			class_id,
			lecturer_id,
			room_id,
			lecture_plan_date,
			lecture_plan_start_time,
			lecture_plan_end_time,
			lecture_actual_date,
			lecture_actual_start_time,
			lecture_actual_end_time,
			is_manual_participation,
			autonomous_participation_start_time,
			autonomous_participation_end_time,
			is_exam,
			is_theory_exam,
			is_practicum_exam,
			is_field_practicum_exam,
			is_midterm_exam,
			is_endterm_exam
		) VALUES (
			:id,
			:class_id,
			:lecturer_id,
			:room_id,
			:lecture_plan_date,
			:lecture_plan_start_time,
			:lecture_plan_end_time,
			:lecture_plan_date,
			:lecture_plan_start_time,
			:lecture_plan_end_time,
			false,
			TO_TIMESTAMP(CONCAT(CAST(DATE(:lecture_plan_date) AS text), TO_CHAR(CAST(:lecture_plan_start_time AS integer), CAST('fm0000' AS text))), 'YYYY-MM-DDHH24MI'),
			TO_TIMESTAMP(CONCAT(CAST(DATE(:lecture_plan_date) AS text), TO_CHAR(CAST(:lecture_plan_end_time AS integer), CAST('fm0000' AS text))), 'YYYY-MM-DDHH24MI'),
			:is_exam,
			:is_theory_exam,
			:is_practicum_exam,
			:is_field_practicum_exam,
			:is_midterm_exam,
			:is_endterm_exam
		)
	`

	updateQuery = `
		UPDATE lectures SET
			lecturer_id = COALESCE(:lecturer_id, lecturer_id),
			foreign_lecturer_name = COALESCE(:foreign_lecturer_name, foreign_lecturer_name),
			foreign_lecturer_source_instance = COALESCE(:foreign_lecturer_source_instance, foreign_lecturer_source_instance),
			is_original_lecturer = COALESCE(:is_original_lecturer, is_original_lecturer),
			room_id = COALESCE(:room_id, room_id),
			lecture_plan_date = COALESCE(:lecture_plan_date, lecture_plan_date),
			lecture_plan_start_time = COALESCE(:lecture_plan_start_time, lecture_plan_start_time),
			lecture_plan_end_time = COALESCE(:lecture_plan_end_time, lecture_plan_end_time),
			lecture_actual_date = COALESCE(:lecture_plan_date, lecture_plan_date),
			lecture_actual_start_time = COALESCE(:lecture_plan_start_time, lecture_plan_start_time),
			lecture_actual_end_time = COALESCE(:lecture_plan_end_time, lecture_plan_end_time),
			lecture_theme = :lecture_theme,
			lecture_subject = :lecture_subject,
			remarks = :remarks,
			autonomous_participation_start_time = TO_TIMESTAMP(CONCAT(CAST(DATE(COALESCE(:lecture_plan_date, lecture_plan_date)) AS text), TO_CHAR(CAST(COALESCE(:lecture_plan_start_time, lecture_plan_start_time) AS integer), CAST('fm0000' AS text))), 'YYYY-MM-DDHH24MI'),
			autonomous_participation_end_time = TO_TIMESTAMP(CONCAT(CAST(DATE(COALESCE(:lecture_plan_date, lecture_plan_date)) AS text), TO_CHAR(CAST(COALESCE(:lecture_plan_end_time, lecture_plan_end_time) AS integer), CAST('fm0000' AS text))), 'YYYY-MM-DDHH24MI')
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM lectures WHERE id = $1
	`

	resetParticipationQuery = `
		UPDATE lecture_participants SET
			is_attend = NULL,
			is_sick = NULL,
			is_leave = NULL,
			is_awol = NULL
		WHERE lecture_id = $1
	`

	getParticipantByLectureIdQuery = `
		SELECT
			lp.id,
			lp.lecture_id,
			lp.student_id,
			lp.is_attend,
			lp.is_sick,
			lp.is_leave,
			lp.is_awol
		FROM lecture_participants lp
		WHERE lp.lecture_id = $1
	`

	bulkUpdateParticipantQuery = `
		INSERT INTO lecture_participants (
			lecture_id,
			student_id,
			is_attend,
			is_sick,
			is_leave,
			is_awol
		) VALUES (
			:lecture_id,
			:student_id,
			:is_attend,
			:is_sick,
			:is_leave,
			:is_awol
		) ON CONFLICT (lecture_id, student_id) DO UPDATE SET
			is_attend = EXCLUDED.is_attend,
			is_sick = EXCLUDED.is_sick,
			is_leave = EXCLUDED.is_leave,
			is_awol = EXCLUDED.is_awol
	`

	getStudentParticipationQuery = `
		SELECT
			l.id,
			l.lecture_plan_date,
			l.lecture_plan_day_of_week,
			l.lecture_plan_start_time,
			l.lecture_plan_end_time,
			l.lecture_actual_date,
			l.lecture_actual_day_of_week,
			l.lecture_actual_start_time,
			l.lecture_actual_end_time,
			lp.is_attend,
			lp.is_sick,
			lp.is_leave,
			lp.is_awol
		FROM lectures l
		JOIN lecture_participants lp ON lp.lecture_id = l.id AND lp.student_id = $2
		WHERE l.class_id = $1
	`

	countStudentParticipationQuery = `
		SELECT COUNT(1)
		FROM lectures l
		JOIN lecture_participants lp ON lp.lecture_id = l.id AND lp.student_id = $2
		WHERE l.class_id = $1
	`

	attendLectureQuery = `
		UPDATE lecture_participants SET is_attend = true
		WHERE lecture_id = $1 AND student_id = $2
	`

	getLectureCalendarQuery = `
		WITH d AS (SELECT generate_series($1::date, $2::date, '1 day'::interval) AS date)
		SELECT 
			d.date,
			l.id AS lecture_id,
			c.id AS class_id,
			c.name AS class_name,
			r.id AS room_id,
			r.name AS room_name,
			lr.id AS lecturer_id,
			lr.name AS lecturer_name,
			lr.front_title AS lecturer_front_title,
			lr.back_degree AS lecturer_back_degree,
			l.lecture_plan_start_time,
			l.lecture_plan_end_time,
			l.foreign_lecturer_name,
			l.foreign_lecturer_source_instance
		FROM d
		LEFT JOIN lectures l ON l.lecture_plan_date = d.date
		LEFT JOIN classes c ON c.id = l.class_id %s
		LEFT JOIN rooms r ON r.id = l.room_id %s
		LEFT JOIN lecturers lr ON lr.id = l.lecturer_id %s
		ORDER BY d.date, l.lecture_plan_start_time
	`

	getHistoryQuery = `
		SELECT
			lp.id,
			l.lecture_actual_date AS lecture_date,
			s.name AS subject_name,
			lp.updated_at AS attend_time
		FROM lecture_participants lp
		JOIN lectures l ON l.id = lp.lecture_id
		JOIN classes c ON c.id = l.class_id
		JOIN subjects s ON s.id = c.subject_id
	`
	countHistoryQuery = `
	SELECT COUNT(1)
	FROM lecture_participants lp
	JOIN lectures l ON l.id = lp.lecture_id
	JOIN classes c ON c.id = l.class_id
	JOIN subjects s ON s.id = c.subject_id
	`
)
