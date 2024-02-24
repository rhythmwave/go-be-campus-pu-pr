package thesis

const (
	getListQuery = `
		SELECT
			t.id,
			t.topic,
			t.title,
			t.status,
			s.id AS student_id,
			s.name AS student_name,
			s.nim_number AS student_nim_number,
			s.status AS student_status,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.code AS dikti_study_program_code,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			s.has_thesis_study_plan AS student_has_thesis_study_plan,
			sse.id AS start_semester_id,
			sse.semester_type AS start_semester_type,
			sse.semester_start_year AS start_semester_start_year
		FROM theses t
		JOIN students s ON s.id = t.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters sse ON sse.id = t.start_semester_id
		%s
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM theses t
		JOIN students s ON s.id = t.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters sse ON sse.id = t.start_semester_id
		%s
	`

	getByIdQuery = `
		SELECT
			t.id,
			s.study_program_id,
			s.id AS student_id,
			s.name AS student_name,
			s.nim_number AS student_nim_number,
			sse.id AS start_semester_id,
			sse.semester_type AS start_semester_type,
			sse.semester_start_year AS start_semester_start_year,
			fse.id AS finish_semester_id,
			fse.semester_type AS finish_semester_type,
			fse.semester_start_year AS finish_semester_start_year,
			t.topic,
			t.title,
			t.english_title,
			t.start_date,
			t.finish_date,
			t.remarks,
			t.is_joint_thesis,
			t.status,
			t.proposal_seminar_date,
			t.proposal_certificate_number,
			t.proposal_certificate_date,
			t.thesis_defense_count,
			t.grade_point,
			t.grade_code
		FROM theses t
		JOIN students s ON s.id = t.student_id
		JOIN semesters sse ON sse.id = t.start_semester_id
		LEFT JOIN semesters fse ON fse.id = t.finish_semester_id
		WHERE t.id = $1
	`

	getByStudentIdStatusQuery = `
		SELECT
			t.id,
			s.study_program_id,
			s.id AS student_id,
			s.name AS student_name,
			s.nim_number AS student_nim_number,
			sse.id AS start_semester_id,
			sse.semester_type AS start_semester_type,
			sse.semester_start_year AS start_semester_start_year,
			fse.id AS finish_semester_id,
			fse.semester_type AS finish_semester_type,
			fse.semester_start_year AS finish_semester_start_year,
			t.topic,
			t.title,
			t.english_title,
			t.start_date,
			t.finish_date,
			t.remarks,
			t.is_joint_thesis,
			t.status,
			t.proposal_seminar_date,
			t.proposal_certificate_number,
			t.proposal_certificate_date,
			t.thesis_defense_count,
			t.grade_point,
			t.grade_code
		FROM theses t
		JOIN students s ON s.id = t.student_id
		JOIN semesters sse ON sse.id = t.start_semester_id
		LEFT JOIN semesters fse ON fse.id = t.finish_semester_id
		WHERE t.student_id = $1 AND t.status = $2
	`

	getNonCancelledQuery = `
		SELECT
			t.id,
			s.study_program_id,
			s.id AS student_id,
			s.name AS student_name,
			s.nim_number AS student_nim_number,
			sse.id AS start_semester_id,
			sse.semester_type AS start_semester_type,
			sse.semester_start_year AS start_semester_start_year,
			fse.id AS finish_semester_id,
			fse.semester_type AS finish_semester_type,
			fse.semester_start_year AS finish_semester_start_year,
			t.topic,
			t.title,
			t.english_title,
			t.start_date,
			t.finish_date,
			t.remarks,
			t.is_joint_thesis,
			t.status,
			t.proposal_seminar_date,
			t.proposal_certificate_number,
			t.proposal_certificate_date,
			t.thesis_defense_count,
			t.grade_point,
			t.grade_code
		FROM theses t
		JOIN students s ON s.id = t.student_id
		JOIN semesters sse ON sse.id = t.start_semester_id
		LEFT JOIN semesters fse ON fse.id = t.finish_semester_id
		WHERE t.student_id = $1 AND t.status != $2
	`

	getFileByThesisIdQuery = `
		SELECT 
			tf.id,
			tf.thesis_id,
			tf.file_path,
			tf.file_path_type,
			tf.description
		FROM thesis_files tf
		WHERE tf.thesis_id = $1
	`

	getSupervisorByThesisIdQuery = `
		SELECT 
			ts.id,
			ts.thesis_id,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			tsr.id AS thesis_supervisor_role_id,
			tsr.name AS thesis_supervisor_role_name,
			tsr.sort AS thesis_supervisor_role_sort
		FROM thesis_supervisors ts
		JOIN lecturers l ON l.id = ts.lecturer_id
		JOIN thesis_supervisor_roles tsr ON tsr.id = ts.thesis_supervisor_role_id
		WHERE ts.thesis_id = $1
	`

	createQuery = `
		INSERT INTO theses (
			id,
			student_id,
			topic,
			status,
			title,
			english_title,
			start_semester_id,
			start_date,
			remarks,
			is_joint_thesis,
			proposal_seminar_date,
			proposal_certificate_number,
			proposal_certificate_date
		) VALUES (
			:id,
			:student_id,
			:topic,
			:status,
			:title,
			:english_title,
			:start_semester_id,
			:start_date,
			:remarks,
			:is_joint_thesis,
			:proposal_seminar_date,
			:proposal_certificate_number,
			:proposal_certificate_date
		)
	`

	upsertFileQuery = `
		INSERT INTO thesis_files (
			thesis_id,
			file_path,
			file_path_type,
			description
		) VALUES (
			:thesis_id,
			:file_path,
			:file_path_type,
			:description
		) ON CONFLICT (thesis_id, file_path) DO UPDATE SET
			description = EXCLUDED.description
	`

	upsertSupervisorQuery = `
		INSERT INTO thesis_supervisors (
			thesis_id,
			lecturer_id,
			thesis_supervisor_role_id
		) VALUES (
			:thesis_id,
			:lecturer_id,
			:thesis_supervisor_role_id
		) ON CONFLICT (thesis_id, lecturer_id) DO UPDATE SET 
			thesis_supervisor_role_id = EXCLUDED.thesis_supervisor_role_id
	`

	deleteFileExcludingPathsQuery = `
		DELETE FROM thesis_files WHERE thesis_id = $1 %s
	`

	deleteSupervisorExcludingLecturerIdQuery = `
		DELETE FROM thesis_supervisors WHERE thesis_id = $1 %s
	`

	updateQuery = `
		UPDATE theses SET
			student_id = :student_id,
			topic = :topic,
			status = :status,
			title = :title,
			english_title = :english_title,
			start_semester_id = :start_semester_id,
			start_date = :start_date,
			remarks = :remarks,
			is_joint_thesis = :is_joint_thesis,
			proposal_seminar_date = :proposal_seminar_date,
			proposal_certificate_number = :proposal_certificate_number,
			proposal_certificate_date = :proposal_certificate_date
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM theses WHERE id = $1
	`

	createDefenseRequestQuery = `
		INSERT INTO thesis_defense_requests (thesis_id) VALUES ($1)
		ON CONFLICT (thesis_id) DO UPDATE SET thesis_defense_id = NULL;
	`

	getListDefenseRequestQuery = `
		SELECT 
			tdr.id,
			s.id AS student_id,
			s.name AS student_name,
			s.nim_number AS student_nim_number,
			s.status AS student_status,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.code AS dikti_study_program_code,
			dsp.type AS dikti_study_program_type,
			sl.id AS study_level_id,
			sl.short_name AS study_level_short_name,
			t.id AS thesis_id,
			t.title AS thesis_title,
			t.status AS thesis_status,
			t.grade_code AS thesis_grade_code,
			tdr.thesis_defense_count,
			td.id AS thesis_defense_id,
			td.plan_date AS thesis_defense_plan_date,
			td.plan_start_time AS thesis_defense_plan_start_time,
			td.plan_end_time AS thesis_defense_plan_end_time,
			td.actual_date AS thesis_defense_actual_date,
			td.actual_start_time AS thesis_defense_actual_start_time,
			td.actual_end_time AS thesis_defense_actual_end_time,
			td.revision AS thesis_defense_revision,
			tdro.id AS thesis_defense_room_id,
			tdro.name AS thesis_defense_room_name,
			td.is_passed AS thesis_defense_is_passed,
			tdr.created_at
		FROM thesis_defense_requests tdr
		JOIN theses t ON t.id = tdr.thesis_id
		LEFT JOIN thesis_defenses td ON td.id = tdr.thesis_defense_id
		LEFT JOIN rooms tdro ON tdro.id = td.room_id
		JOIN students s ON s.id = t.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters sse ON sse.id = t.start_semester_id
	`

	countListDefenseRequestQuery = `
		SELECT COUNT(1)
		FROM thesis_defense_requests tdr
		JOIN theses t ON t.id = tdr.thesis_id
		LEFT JOIN thesis_defenses td ON td.id = tdr.thesis_defense_id
		JOIN students s ON s.id = t.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters sse ON sse.id = t.start_semester_id
	`

	getActiveDefenseRequestQuery = `
		SELECT 
			tdr.id,
			s.id AS student_id,
			s.name AS student_name,
			s.nim_number AS student_nim_number,
			s.status AS student_status,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.code AS dikti_study_program_code,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			t.id AS thesis_id,
			t.title AS thesis_title,
			t.status AS thesis_status,
			tdr.thesis_defense_count,
			td.id AS thesis_defense_id,
			td.plan_date AS thesis_defense_plan_date,
			td.plan_start_time AS thesis_defense_plan_start_time,
			td.plan_end_time AS thesis_defense_plan_end_time,
			td.actual_date AS thesis_defense_actual_date,
			td.actual_start_time AS thesis_defense_actual_start_time,
			td.actual_end_time AS thesis_defense_actual_end_time,
			tdro.id AS thesis_defense_room_id,
			tdro.name AS thesis_defense_room_name,
			td.is_passed AS thesis_defense_is_passed,
			tdr.created_at
		FROM thesis_defense_requests tdr
		JOIN theses t ON t.id = tdr.thesis_id
		LEFT JOIN thesis_defenses td ON td.id = tdr.thesis_defense_id
		LEFT JOIN rooms tdro ON tdro.id = td.room_id
		JOIN students s ON s.id = t.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters sse ON sse.id = t.start_semester_id
		WHERE t.id = $1 AND (td.id IS NULL OR td.actual_date IS NULL OR td.is_passed IS true)
	`

	createDefenseQuery = `
		INSERT INTO thesis_defenses (
			id,
			thesis_id,
			plan_date,
			plan_start_time,
			plan_end_time,
			room_id
		) VALUES (
			:id,
			:thesis_id,
			:plan_date,
			:plan_start_time,
			:plan_end_time,
			:room_id
		)
	`

	deleteExaminerExcludingLecturerIdQuery = `
		DELETE FROM thesis_examiners WHERE thesis_defense_id = $1 %s
	`

	upsertDefenseExaminerQuery = `
		INSERT INTO thesis_examiners (
			thesis_defense_id,
			lecturer_id,
			thesis_examiner_role_id
		) VALUES (
			:thesis_defense_id,
			:lecturer_id,
			:thesis_examiner_role_id
		) ON CONFLICT (thesis_defense_id, lecturer_id) DO UPDATE SET thesis_examiner_role_id = EXCLUDED.thesis_examiner_role_id
	`

	getDefenseByIdQuery = `
		SELECT
			td.id,
			td.plan_date,
			td.plan_start_time,
			td.plan_end_time,
			td.actual_date,
			td.actual_start_time,
			td.actual_end_time,
			td.is_passed,
			r.id AS room_id,
			r.name AS room_name,
			s.id AS student_id,
			s.name AS student_name,
			s.nim_number AS student_nim_number,
			s.status AS student_status,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			dsp.code AS dikti_study_program_code,
			dsp.type AS dikti_study_program_type,
			sl.id AS study_level_id,
			sl.short_name AS study_level_short_name,
			t.id AS thesis_id,
			t.title AS thesis_title,
			t.status AS thesis_status,
			td.thesis_defense_count,
			td.created_at
		FROM thesis_defenses td
		JOIN theses t ON t.id = td.thesis_id
		JOIN rooms r ON r.id = td.room_id
		JOIN students s ON s.id = t.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		JOIN study_levels sl ON sl.id = dsp.study_level_id
		JOIN semesters sse ON sse.id = t.start_semester_id
		WHERE td.id = $1
	`

	updateDefenseQuery = `
		UPDATE thesis_defenses SET
			plan_date = :plan_date,
			plan_start_time = :plan_start_time,
			plan_end_time = :plan_end_time,
			room_id = :room_id,
			actual_date = :actual_date,
			actual_start_time = :actual_start_time,
			actual_end_time = :actual_end_time,
			is_passed = :is_passed,
			revision = :revision
		WHERE id = :id
	`

	finishDefenseQuery = `
		UPDATE theses SET
			finish_semester_id = (CASE WHEN finish_semester_id IS NOT NULL THEN finish_semester_id ELSE :finish_semester_id END),
			finish_date = (CASE WHEN finish_date IS NOT NULL THEN finish_date ELSE now() END),
			grade_point = :grade_point,
			grade_code = :grade_code,
			status = :status
		WHERE id = :id
	`

	getThesisDefenseExaminerByThesisDefenseIdsQuery = `
		SELECT
			te.id,
			te.thesis_defense_id,
			l.id as lecturer_id,
			l.name as lecturer_name,
			l.front_title as lecturer_front_title,
			l.back_degree as lecturer_back_degree,
			te.thesis_examiner_role_id,
			ter.name AS thesis_examiner_role_name
		FROM thesis_examiners te
		JOIN lecturers l ON l.id = te.lecturer_id
		JOIN thesis_examiner_roles ter ON ter.id = te.thesis_examiner_role_id
		WHERE te.thesis_defense_id IN (SELECT UNNEST($1::uuid[]))
	`

	getActiveSemesterThesisSupervisorLogQuery = `
		SELECT
			l.id AS lecturer_id,
			tsr.id AS thesis_supervisor_role_id,
			tsr.name AS thesis_supervisor_role_name,
			COUNT(t.id) AS total
		FROM lecturers l
		CROSS JOIN thesis_supervisor_roles tsr
		LEFT JOIN thesis_supervisors ts ON ts.lecturer_id = l.id AND ts.thesis_supervisor_role_id = tsr.id
		LEFT JOIN theses t ON t.id = ts.thesis_id AND t.status::text = $2
		WHERE l.id IN (SELECT UNNEST($1::uuid[]))
		GROUP BY l.id, tsr.id
	`

	getThesisSupervisorLogQuery = `
		SELECT
			l.id AS lecturer_id,
			tsr.id AS thesis_supervisor_role_id,
			tsr.name AS thesis_supervisor_role_name,
			COALESCE(tsl.total, 0) AS total
		FROM lecturers l
		CROSS JOIN thesis_supervisor_roles tsr
		LEFT JOIN thesis_supervisor_logs tsl ON tsl.lecturer_id = l.id AND tsl.semester_id = $1 AND tsl.thesis_supervisor_role_id	= tsr.id
		WHERE l.id IN (SELECT UNNEST($2::uuid[]))
	`
)
