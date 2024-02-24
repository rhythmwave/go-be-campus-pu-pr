package exam_supervisor

const (
	getListQuery = `
		SELECT 
			es.id, 
			es.name,
			es.id_national_lecturer
		FROM exam_supervisors es
		LEFT JOIN study_programs sp ON sp.id = es.study_program_id
	`

	countListQuery = `
		SELECT COUNT(1) FROM exam_supervisors es
		LEFT JOIN study_programs sp ON sp.id = es.study_program_id
	`

	getDetailQuery = `
		SELECT
			es.id,
			es.id_national_lecturer,
			es.name,
			es.front_title,
			es.back_degree,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			es.id_number,
			es.birth_date,
			br.id AS birth_regency_id,
			br.name AS birth_regency_name,
			bc.id AS birth_country_id,
			bc.name AS birth_country_name,
			es.sex,
			es.blood_type,
			es.religion,
			es.marital_status,
			es.address,
			r.id AS regency_id,
			r.name AS regency_name,
			c.id AS country_id,
			c.name AS country_name,
			es.postal_code,
			es.phone_number,
			es.fax,
			es.mobile_phone_number,
			es.office_phone_number,
			es.employee_type,
			es.employee_status,
			es.sk_cpns_number,
			es.sk_cpns_date,
			es.tmt_cpns_date,
			es.cpns_category,
			es.cpns_duration_month,
			es.pre_position_date,
			es.sk_pns_number,
			es.sk_pns_date,
			es.tmt_pns_date,
			es.pns_category,
			es.pns_oath_date,
			es.join_date,
			es.end_date,
			es.taspen_number,
			es.former_instance,
			es.remarks
		FROM exam_supervisors es
		LEFT JOIN study_programs sp ON sp.id = es.study_program_id
		LEFT JOIN regencies br ON br.id = es.birth_regency_id
		LEFT JOIN provinces bp ON bp.id = br.province_id
		LEFT JOIN countries bc ON bc.id = bp.country_id
		LEFT JOIN regencies r ON r.id = es.birth_regency_id
		LEFT JOIN provinces p ON p.id = r.province_id
		LEFT JOIN countries c ON c.id = p.country_id
		WHERE es.id = $1
	`

	getDetailByIdsQuery = `
		SELECT
			es.id,
			es.id_national_lecturer,
			es.name,
			es.front_title,
			es.back_degree,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			es.id_number,
			es.birth_date,
			br.id AS birth_regency_id,
			br.name AS birth_regency_name,
			bc.id AS birth_country_id,
			bc.name AS birth_country_name,
			es.sex,
			es.blood_type,
			es.religion,
			es.marital_status,
			es.address,
			r.id AS regency_id,
			r.name AS regency_name,
			c.id AS country_id,
			c.name AS country_name,
			es.postal_code,
			es.phone_number,
			es.fax,
			es.mobile_phone_number,
			es.office_phone_number,
			es.employee_type,
			es.employee_status,
			es.sk_cpns_number,
			es.sk_cpns_date,
			es.tmt_cpns_date,
			es.cpns_category,
			es.cpns_duration_month,
			es.pre_position_date,
			es.sk_pns_number,
			es.sk_pns_date,
			es.tmt_pns_date,
			es.pns_category,
			es.pns_oath_date,
			es.join_date,
			es.end_date,
			es.taspen_number,
			es.former_instance,
			es.remarks
		FROM exam_supervisors es
		LEFT JOIN study_programs sp ON sp.id = es.study_program_id
		LEFT JOIN regencies br ON br.id = es.birth_regency_id
		LEFT JOIN provinces bp ON bp.id = br.province_id
		LEFT JOIN countries bc ON bc.id = bp.country_id
		LEFT JOIN regencies r ON r.id = es.birth_regency_id
		LEFT JOIN provinces p ON p.id = r.province_id
		LEFT JOIN countries c ON c.id = p.country_id
		WHERE es.id IN (SELECT UNNEST($1::uuid[]))
	`

	createQuery = `
		INSERT INTO exam_supervisors (
			id_national_lecturer,
			name,
			front_title,
			back_degree,
			study_program_id,
			id_number,
			birth_date,
			birth_regency_id,
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
			:created_by
		)
	`

	updateQuery = `
		UPDATE exam_supervisors SET
			id_national_lecturer = :id_national_lecturer,
			name = :name,
			front_title = :front_title,
			back_degree = :back_degree,
			study_program_id = :study_program_id,
			id_number = :id_number,
			birth_date = :birth_date,
			birth_regency_id = :birth_regency_id,
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
			updated_by = :updated_by
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM exam_supervisors WHERE id = $1
	`

	getExamLectureSupervisorByLectureIdsQuery = `
		SELECT
			ess.lecture_id,
			es.id,
			es.id_national_lecturer,
			es.name,
			es.front_title,
			es.back_degree,
			esr.id AS exam_supervisor_role_id,
			esr.name AS exam_supervisor_role_name
		FROM exam_lecture_supervisors ess
		JOIN exam_supervisors es ON es.id = ess.exam_supervisor_id
		JOIN exam_supervisor_roles esr ON esr.id = ess.exam_supervisor_role_id
		WHERE ess.lecture_id IN (SELECT UNNEST($1::uuid[]))
	`

	deleteExamLectureSupervisorExcludingExamSupervisorIdsQuery = `
		DELETE FROM exam_lecture_supervisors WHERE lecture_id = $1 %s
	`

	upsertExamLectureSupervisorQuery = `
		INSERT INTO exam_lecture_supervisors (
			lecture_id,
			exam_supervisor_id,
			exam_supervisor_role_id,
			created_by
		) VALUES (
			:lecture_id,
			:exam_supervisor_id,
			:exam_supervisor_role_id,
			:created_by
		) ON CONFLICT (lecture_id, exam_supervisor_id) DO UPDATE SET
			exam_supervisor_role_id	= EXCLUDED.exam_supervisor_role_id,
			updated_by = EXCLUDED.created_by
	`
)
