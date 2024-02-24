package student_activity

const (
	getListQuery = `
		SELECT
			sa.id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			sa.semester_id,
			s.semester_start_year,
			s.semester_type,
			sa.activity_type,
			sa.title
		FROM student_activities sa
		JOIN study_programs sp ON sp.id = sa.study_program_id
		JOIN semesters s ON s.id = sa.semester_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM student_activities sa
		JOIN study_programs sp ON sp.id = sa.study_program_id
		JOIN semesters s ON s.id = sa.semester_id
	`

	getListParticipantByStudentActivityIdQuery = `
		SELECT
			sap.student_id,
			s.nim_number,
			s.name,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			sap.role
		FROM student_activity_participants sap
		JOIN students s ON s.id = sap.student_id
		LEFT JOIN study_programs sp ON sp.id = s.study_program_id
		WHERE sap.student_activity_id = $1
		ORDER BY s.name
	`

	getListLecturerByStudentActivityIdQuery = `
		SELECT
			sal.lecturer_id,
			l.id_national_lecturer,
			l.name,
			l.front_title,
			l.back_degree,
			sal.activity_category,
			sal.role,
			sal.sort
		FROM student_activity_lecturers sal
		JOIN lecturers l ON l.id = sal.lecturer_id
		WHERE sal.student_activity_id = $1
		ORDER BY sal.role, sal.sort, l.name
	`

	getDetailQuery = `
		SELECT 
			sa.id,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			sa.semester_id,
			s.semester_start_year,
			s.semester_type,
			sa.activity_type,
			sa.title,
			sa.location,
			sa.decision_number,
			sa.decision_date,
			sa.is_group_activity,
			sa.remarks
		FROM student_activities sa
		JOIN study_programs sp ON sp.id = sa.study_program_id
		JOIN semesters s ON s.id = sa.semester_id
		WHERE sa.id = $1
	`

	createQuery = `
		INSERT INTO student_activities (
			study_program_id,
			semester_id,
			activity_type,
			title,
			location,
			decision_number,
			decision_date,
			is_group_activity,
			remarks,
			is_mbkm,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id;
	`

	deleteParticipantExcludingStudentIdsQuery = `
		DELETE FROM student_activity_participants
		WHERE student_activity_id = $1 %s
	`

	deleteLecturerExcludingLecturerIdsQuery = `
		DELETE FROM student_activity_lecturers
		WHERE student_activity_id = $1 AND role = $2 %s
	`

	upsertParticipantQuery = `
		INSERT INTO student_activity_participants (
			student_activity_id,
			student_id,
			role,
			created_by
		) VALUES (
			:student_activity_id,
			:student_id,
			:role,
			:created_by
		) ON CONFLICT (student_activity_id, student_id) DO UPDATE SET
			role = EXCLUDED.role,
			updated_by = EXCLUDED.created_by
	`

	upsertLecturerQuery = `
		INSERT INTO student_activity_lecturers (
			student_activity_id,
			lecturer_id,
			activity_category,
			role,
			sort,
			created_by
		) VALUES (
			:student_activity_id,
			:lecturer_id,
			:activity_category,
			:role,
			:sort,
			:created_by
		) ON CONFLICT (student_activity_id, lecturer_id, role) DO UPDATE SET
			activity_category = EXCLUDED.activity_category,
			sort = EXCLUDED.sort,
			updated_by = EXCLUDED.created_by
	`

	updateQuery = `
		UPDATE student_activities SET
			study_program_id = :study_program_id,
			semester_id = :semester_id,
			activity_type = :activity_type,
			title = :title,
			location = :location,
			decision_number = :decision_number,
			decision_date = :decision_date,
			is_group_activity = :is_group_activity,
			remarks = :remarks,
			is_mbkm = :is_mbkm,
			updated_by = :updated_by
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM student_activities WHERE id = $1
	`
)
