package lecturer

const (
	getListQuery = `
		SELECT 
			l.id, 
			l.name,
			l.phone_number,
			l.mobile_phone_number,
			l.office_phone_number,
			l.id_national_lecturer,
			l.front_title,
			l.back_degree,
			dsp.code AS dikti_study_program_code,
			sp.name AS study_program_name,
			l.employment_status,
			l.status,
			a.id AS authentication_id,
			a.is_active AS authentication_is_active,
			a.suspension_remarks AS authentication_suspension_remarks,
			l.total_supervised_thesis
			%s
		FROM lecturers l
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN authentications a ON a.lecturer_id = l.id
		%s
	`

	countListQuery = `
		SELECT COUNT(1) FROM lecturers l
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN authentications a ON a.lecturer_id = l.id
		%s
	`

	getScheduleQuery = `
		SELECT
			l.id,
			lr.id_national_lecturer AS id_national_lecturer,
			lr.name AS name,
			lr.front_title AS front_title,
			lr.back_degree AS back_degree,
			sp.name AS study_program_name,
			su.name AS subject_name,
			c.name AS class_name,
			(su.theory_credit + su.practicum_credit + su.field_practicum_credit) AS total_subject_credit,
			l.lecture_plan_date,
			l.lecture_plan_start_time AS start_time,
			l.lecture_plan_end_time AS end_time,
			r.name AS room_name,
			c.total_participant
		FROM lectures l
		JOIN rooms r ON r.id = l.room_id
		JOIN classes c ON c.id = l.class_id
		JOIN subjects su ON su.id = c.subject_id
		JOIN curriculums cu ON cu.id = su.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN lecturers lr ON lr.id = l.lecturer_id
	`

	countScheduleQuery = `
		SELECT COUNT(1)
		FROM lectures l
		JOIN rooms r ON r.id = l.room_id
		JOIN classes c ON c.id = l.class_id
		JOIN subjects su ON su.id = c.subject_id
		JOIN curriculums cu ON cu.id = su.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN lecturers lr ON lr.id = l.lecturer_id
	`

	getDetailQuery = `
		SELECT
			l.id,
			l.id_national_lecturer,
			l.name,
			l.front_title,
			l.back_degree,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			l.id_number,
			l.birth_date,
			br.id AS birth_regency_id,
			br.name AS birth_regency_name,
			bc.id AS birth_country_id,
			bc.name AS birth_country_name,
			l.id_employee,
			l.stambuk,
			l.sex,
			l.blood_type,
			l.religion,
			l.marital_status,
			l.address,
			r.id AS regency_id,
			r.name AS regency_name,
			c.id AS country_id,
			c.name AS country_name,
			l.postal_code,
			l.phone_number,
			l.fax,
			l.mobile_phone_number,
			l.office_phone_number,
			l.employee_type,
			l.employee_status,
			l.sk_cpns_number,
			l.sk_cpns_date,
			l.tmt_cpns_date,
			l.cpns_category,
			l.cpns_duration_month,
			l.pre_position_date,
			l.sk_pns_number,
			l.sk_pns_date,
			l.tmt_pns_date,
			l.pns_category,
			l.pns_oath_date,
			l.join_date,
			l.end_date,
			l.taspen_number,
			l.former_instance,
			l.remarks,
			l.lecturer_number,
			l.academic_position,
			l.employment_status,
			l.expertise,
			l.highest_degree,
			l.instance_code,
			l.teaching_certificate_number,
			l.teaching_permit_number,
			l.status,
			l.resign_semester,
			eg.id AS expertise_group_id,
			eg.name AS expertise_group_name
		FROM lecturers l
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN regencies br ON br.id = l.birth_regency_id
		LEFT JOIN provinces bp ON bp.id = br.province_id
		LEFT JOIN countries bc ON bc.id = bp.country_id
		LEFT JOIN regencies r ON r.id = l.birth_regency_id
		LEFT JOIN provinces p ON p.id = r.province_id
		LEFT JOIN countries c ON c.id = p.country_id
		LEFT JOIN expertise_groups eg ON eg.id = l.expertise_group_id
		WHERE l.id = $1
	`

	getDetailByIdsQuery = `
		SELECT
			l.id,
			l.id_national_lecturer,
			l.name,
			l.front_title,
			l.back_degree,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			l.id_number,
			l.birth_date,
			br.id AS birth_regency_id,
			br.name AS birth_regency_name,
			bc.id AS birth_country_id,
			bc.name AS birth_country_name,
			l.id_employee,
			l.stambuk,
			l.sex,
			l.blood_type,
			l.religion,
			l.marital_status,
			l.address,
			r.id AS regency_id,
			r.name AS regency_name,
			c.id AS country_id,
			c.name AS country_name,
			l.postal_code,
			l.phone_number,
			l.fax,
			l.mobile_phone_number,
			l.office_phone_number,
			l.employee_type,
			l.employee_status,
			l.sk_cpns_number,
			l.sk_cpns_date,
			l.tmt_cpns_date,
			l.cpns_category,
			l.cpns_duration_month,
			l.pre_position_date,
			l.sk_pns_number,
			l.sk_pns_date,
			l.tmt_pns_date,
			l.pns_category,
			l.pns_oath_date,
			l.join_date,
			l.end_date,
			l.taspen_number,
			l.former_instance,
			l.remarks,
			l.lecturer_number,
			l.academic_position,
			l.employment_status,
			l.expertise,
			l.highest_degree,
			l.instance_code,
			l.teaching_certificate_number,
			l.teaching_permit_number,
			l.status,
			l.resign_semester,
			eg.id AS expertise_group_id,
			eg.name AS expertise_group_name
		FROM lecturers l
		LEFT JOIN study_programs sp ON sp.id = l.study_program_id
		LEFT JOIN regencies br ON br.id = l.birth_regency_id
		LEFT JOIN provinces bp ON bp.id = br.province_id
		LEFT JOIN countries bc ON bc.id = bp.country_id
		LEFT JOIN regencies r ON r.id = l.birth_regency_id
		LEFT JOIN provinces p ON p.id = r.province_id
		LEFT JOIN countries c ON c.id = p.country_id
		LEFT JOIN expertise_groups eg ON eg.id = l.expertise_group_id
		WHERE l.id IN (SELECT UNNEST($1::uuid[]))
	`

	createQuery = `
		INSERT INTO lecturers (
			id_national_lecturer,
			name,
			front_title,
			back_degree,
			study_program_id,
			id_number,
			birth_date,
			birth_regency_id,
			id_employee,
			stambuk,
			sex,
			blood_type,
			religion,
			marital_status,
			address,
			regency_id,
			postal_code,
			phone_number,
			fax,
			mobile_phone_number,
			office_phone_number,
			employee_type,
			employee_status,
			sk_cpns_number,
			sk_cpns_date,
			tmt_cpns_date,
			cpns_category,
			cpns_duration_month,
			pre_position_date,
			sk_pns_number,
			sk_pns_date,
			tmt_pns_date,
			pns_category,
			pns_oath_date,
			join_date,
			end_date,
			taspen_number,
			former_instance,
			remarks,
			lecturer_number,
			academic_position,
			employment_status,
			expertise,
			highest_degree,
			instance_code,
			teaching_certificate_number,
			teaching_permit_number,
			status,
			resign_semester,
			expertise_group_id,
			created_by
		) VALUES (
			:id_national_lecturer,
			:name,
			:front_title,
			:back_degree,
			:study_program_id,
			:id_number,
			:birth_date,
			:birth_regency_id,
			:id_employee,
			:stambuk,
			:sex,
			:blood_type,
			:religion,
			:marital_status,
			:address,
			:regency_id,
			:postal_code,
			:phone_number,
			:fax,
			:mobile_phone_number,
			:office_phone_number,
			:employee_type,
			:employee_status,
			:sk_cpns_number,
			:sk_cpns_date,
			:tmt_cpns_date,
			:cpns_category,
			:cpns_duration_month,
			:pre_position_date,
			:sk_pns_number,
			:sk_pns_date,
			:tmt_pns_date,
			:pns_category,
			:pns_oath_date,
			:join_date,
			:end_date,
			:taspen_number,
			:former_instance,
			:remarks,
			:lecturer_number,
			:academic_position,
			:employment_status,
			:expertise,
			:highest_degree,
			:instance_code,
			:teaching_certificate_number,
			:teaching_permit_number,
			:status,
			:resign_semester,
			:expertise_group_id,
			:created_by
		)
	`

	updateQuery = `
		UPDATE lecturers SET
			id_national_lecturer = :id_national_lecturer,
			name = :name,
			front_title = :front_title,
			back_degree = :back_degree,
			study_program_id = :study_program_id,
			id_number = :id_number,
			birth_date = :birth_date,
			birth_regency_id = :birth_regency_id,
			id_employee = :id_employee,
			stambuk = :stambuk,
			sex = :sex,
			blood_type = :blood_type,
			religion = :religion,
			marital_status = :marital_status,
			address = :address,
			regency_id = :regency_id,
			postal_code = :postal_code,
			phone_number = :phone_number,
			fax = :fax,
			mobile_phone_number = :mobile_phone_number,
			office_phone_number = :office_phone_number,
			employee_type = :employee_type,
			employee_status = :employee_status,
			sk_cpns_number = :sk_cpns_number,
			sk_cpns_date = :sk_cpns_date,
			tmt_cpns_date = :tmt_cpns_date,
			cpns_category = :cpns_category,
			cpns_duration_month = :cpns_duration_month,
			pre_position_date = :pre_position_date,
			sk_pns_number = :sk_pns_number,
			sk_pns_date = :sk_pns_date,
			tmt_pns_date = :tmt_pns_date,
			pns_category = :pns_category,
			pns_oath_date = :pns_oath_date,
			join_date = :join_date,
			end_date = :end_date,
			taspen_number = :taspen_number,
			former_instance = :former_instance,
			remarks = :remarks,
			lecturer_number = :lecturer_number,
			academic_position = :academic_position,
			employment_status = :employment_status,
			expertise = :expertise,
			highest_degree = :highest_degree,
			instance_code = :instance_code,
			teaching_certificate_number = :teaching_certificate_number,
			teaching_permit_number = :teaching_permit_number,
			status = :status,
			resign_semester = :resign_semester,
			expertise_group_id = :expertise_group_id,
			updated_by = :updated_by
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM lecturers WHERE id = $1
	`

	updateStatusQuery = `
		UPDATE lecturers SET status = $1
		WHERE id IN (SELECT UNNEST($2::uuid[]))
	`

	getAssignedClassQuery = `
		SELECT 
			c.id,
			c.name,
			s.code AS subject_code,
			s.name AS subject_name,
			s.theory_credit,
			s.practicum_credit,
			s.field_practicum_credit,
			cl.is_grading_responsible,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			cl.total_attendance,
			c.total_lecture_plan,
			c.total_lecture_done
		FROM classes c
		JOIN subjects s ON s.id = c.subject_id
		JOIN curriculums cu ON cu.id = s.curriculum_id
		JOIN study_programs sp ON sp.id = cu.study_program_id
		JOIN class_lecturers cl ON cl.class_id = c.id AND cl.lecturer_id = $1
		%s
		ORDER BY s.code, c.name
	`
)
