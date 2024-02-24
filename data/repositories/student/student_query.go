package student

const (
	getListQuery = `
		SELECT 
			s.id,
			s.name,
			s.sex,
			s.marital_status,
			br.id AS birth_regency_id,
			br.name AS birth_regency_name,
			s.birth_date,
			s.religion,
			s.address,
			s.rt,
			s.rw,
			v.id AS village_id,
			v.name AS village_name,
			d.id AS district_id,
			d.name AS district_name,
			r.id AS regency_id,
			r.name AS regency_name,
			p.id AS province_id,
			p.name AS province_name,
			c.id AS country_id,
			c.name AS country_name,
			s.postal_code,
			s.previous_address,
			s.id_number,
			s.npwp_number,
			s.nisn_number,
			s.residence_type,
			s.transportation_mean,
			s.kps_receiver,
			s.phone_number,
			s.mobile_phone_number,
			s.email,
			s.homepage,
			s.work_type,
			s.work_place,
			s.nationality,
			s.askes_number,
			s.total_brother,
			s.total_sister,
			s.hobby,
			s.experience,
			s.total_dependent,
			s.nim_number,
			s.student_force,
			s.admittance_semester,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			sp.minimum_graduation_credit AS study_program_minimum_graduation_credit,
			sp.minimum_thesis_credit AS study_program_minimum_thesis_credit,
			dsp.code AS dikti_study_program_code,
			cu.id AS curriculum_id,
			cu.name AS curriculum_name,
			s.admittance_test_number,
			s.admittance_date,
			s.admittance_status,
			s.total_transfer_credit,
			s.previous_college,
			s.previous_study_program,
			s.previous_nim_number,
			s.previous_nim_admittance_year,
			s.status,
			s.is_foreign_student,
			s.college_entrance_type,
			s.class_time,
			s.fund_source,
			s.is_scholarship_grantee,
			s.has_complete_requirement,
			s.school_certificate_type,
			s.school_graduation_year,
			s.school_name,
			s.school_accreditation,
			s.school_address,
			s.school_major,
			s.school_certificate_number,
			s.school_certificate_date,
			s.total_school_final_exam_subject,
			s.school_final_exam_score,
			s.guardian_name,
			s.guardian_birth_date,
			s.guardian_death_date,
			s.guardian_address,
			gr.id AS guardian_regency_id,
			gr.name AS guardian_regency_name,
			s.guardian_postal_code,
			s.guardian_phone_number,
			s.guardian_email,
			s.guardian_final_academic_background,
			s.guardian_occupation,
			s.father_id_number,
			s.father_name,
			s.father_birth_date,
			s.father_death_date,
			s.mother_id_number,
			s.mother_name,
			s.mother_birth_date,
			s.mother_death_date,
			s.parent_address,
			pr.id AS parent_regency_id,
			pr.name AS parent_regency_name,
			s.parent_postal_code,
			s.parent_phone_number,
			s.parent_email,
			s.father_final_academic_background,
			s.father_occupation,
			s.mother_final_academic_background,
			s.mother_occupation,
			s.parent_income,
			s.is_financially_capable,
			a.id AS authentication_id,
			a.is_active AS authentication_is_active,
			a.suspension_remarks AS authentication_suspension_remarks,
			s.current_semester_package,
			s.total_study_plan,
			s.has_paid,
			s.blood_type,
			s.profile_photo_path,
			s.profile_photo_path_type,
			bp.id AS birth_provice_id,
			bp.name AS birth_provice_name,
			s.height,
			s.weight,
			s.is_color_blind,
			s.use_glasses,
			s.has_complete_teeth,
			s.is_kps_recipient,
			s.work_address,
			s.assurance_number,
			dsp.type AS dikti_study_program_type,
			sl.name AS study_level_name,
			sl.short_name AS study_level_short_name,
			s.parent_religion,
			s.parent_nationality,
			s.father_work_address,
			pp.id AS parent_province_id,
			pp.name AS parent_province_name,
			gp.id AS guardian_province_id,
			gp.name AS guardian_province_name,
			s.school_enrollment_year,
			s.school_enrollment_class,
			s.school_type,
			scp.id AS school_province_id,
			scp.name AS school_province_name,
			s.school_status,
			s.school_mathematics_final_exam_score,
			s.school_indonesian_final_exam_score,
			s.school_english_final_exam_score,
			s.school_mathematics_report_score,
			s.school_indonesian_report_score,
			s.school_english_report_score,
			s.gpa,
			s.total_credit,
			s.transcript_is_archived,
			s.graduation_date,
			s.diploma_number,
			grp.predicate AS graduation_predicate,
			t.title AS thesis_title,
			t.english_title AS thesis_english_title,
			t.start_semester_id AS thesis_start_semester_id,
			t.finish_semester_id AS thesis_finish_semester_id,
			t.start_date AS thesis_start_date,
			t.finish_date AS thesis_finish_date,
			s.created_at
			%s
		FROM students s
		LEFT JOIN regencies br ON br.id = s.birth_regency_id
		LEFT JOIN provinces bp ON bp.id = br.province_id
		LEFT JOIN villages v ON v.id = s.village_id
		LEFT JOIN districts d ON d.id = v.district_id
		LEFT JOIN regencies r ON r.id = d.regency_id
		LEFT JOIN provinces p ON p.id = r.province_id
		LEFT JOIN countries c ON c.id = p.country_id
		LEFT JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		LEFT JOIN curriculums cu ON cu.id = s.curriculum_id
		LEFT JOIN regencies gr ON gr.id = s.guardian_regency_id
		LEFT JOIN provinces gp ON gp.id = gr.province_id
		LEFT JOIN regencies pr ON pr.id = s.parent_regency_id
		LEFT JOIN provinces pp ON pp.id = pr.province_id
		LEFT JOIN authentications a ON a.student_id = s.id
		LEFT JOIN provinces scp ON scp.id = s.school_province_id
		LEFT JOIN graduation_predicates grp ON grp.id = s.graduation_predicate_id
		LEFT JOIN theses t ON t.student_id = s.id AND t.status::text = $1
		%s
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM students s
		LEFT JOIN regencies br ON br.id = s.birth_regency_id
		LEFT JOIN provinces bp ON bp.id = br.province_id
		LEFT JOIN villages v ON v.id = s.village_id
		LEFT JOIN districts d ON d.id = v.district_id
		LEFT JOIN regencies r ON r.id = d.regency_id
		LEFT JOIN provinces p ON p.id = r.province_id
		LEFT JOIN countries c ON c.id = p.country_id
		LEFT JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		LEFT JOIN curriculums cu ON cu.id = s.curriculum_id
		LEFT JOIN regencies gr ON gr.id = s.guardian_regency_id
		LEFT JOIN provinces gp ON gp.id = gr.province_id
		LEFT JOIN regencies pr ON pr.id = s.parent_regency_id
		LEFT JOIN provinces pp ON pp.id = pr.province_id
		LEFT JOIN authentications a ON a.student_id = s.id
		LEFT JOIN provinces scp ON scp.id = s.school_province_id
		LEFT JOIN graduation_predicates grp ON grp.id = s.graduation_predicate_id
		LEFT JOIN theses t ON t.student_id = s.id AND t.status::text = $1
		%s
	`

	getDetailQuery = `
		SELECT 
			s.id,
			s.name,
			s.sex,
			s.marital_status,
			br.id AS birth_regency_id,
			br.name AS birth_regency_name,
			s.birth_date,
			s.religion,
			s.address,
			s.rt,
			s.rw,
			v.id AS village_id,
			v.name AS village_name,
			d.id AS district_id,
			d.name AS district_name,
			r.id AS regency_id,
			r.name AS regency_name,
			p.id AS province_id,
			p.name AS province_name,
			c.id AS country_id,
			c.name AS country_name,
			s.postal_code,
			s.previous_address,
			s.id_number,
			s.npwp_number,
			s.nisn_number,
			s.residence_type,
			s.transportation_mean,
			s.kps_receiver,
			s.phone_number,
			s.mobile_phone_number,
			s.email,
			s.homepage,
			s.work_type,
			s.work_place,
			s.nationality,
			s.askes_number,
			s.total_brother,
			s.total_sister,
			s.hobby,
			s.experience,
			s.total_dependent,
			s.nim_number,
			s.student_force,
			s.admittance_semester,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			sp.minimum_graduation_credit AS study_program_minimum_graduation_credit,
			sp.minimum_thesis_credit AS study_program_minimum_thesis_credit,
			dsp.code AS dikti_study_program_code,
			cu.id AS curriculum_id,
			cu.name AS curriculum_name,
			s.admittance_test_number,
			s.admittance_date,
			s.admittance_status,
			s.total_transfer_credit,
			s.previous_college,
			s.previous_study_program,
			s.previous_nim_number,
			s.previous_nim_admittance_year,
			s.status,
			s.is_foreign_student,
			s.college_entrance_type,
			s.class_time,
			s.fund_source,
			s.is_scholarship_grantee,
			s.has_complete_requirement,
			s.school_certificate_type,
			s.school_graduation_year,
			s.school_name,
			s.school_accreditation,
			s.school_address,
			s.school_major,
			s.school_certificate_number,
			s.school_certificate_date,
			s.total_school_final_exam_subject,
			s.school_final_exam_score,
			s.guardian_name,
			s.guardian_birth_date,
			s.guardian_death_date,
			s.guardian_address,
			gr.id AS guardian_regency_id,
			gr.name AS guardian_regency_name,
			s.guardian_postal_code,
			s.guardian_phone_number,
			s.guardian_email,
			s.guardian_final_academic_background,
			s.guardian_occupation,
			s.father_id_number,
			s.father_name,
			s.father_birth_date,
			s.father_death_date,
			s.mother_id_number,
			s.mother_name,
			s.mother_birth_date,
			s.mother_death_date,
			s.parent_address,
			pr.id AS parent_regency_id,
			pr.name AS parent_regency_name,
			s.parent_postal_code,
			s.parent_phone_number,
			s.parent_email,
			s.father_final_academic_background,
			s.father_occupation,
			s.mother_final_academic_background,
			s.mother_occupation,
			s.parent_income,
			s.is_financially_capable,
			a.id AS authentication_id,
			a.is_active AS authentication_is_active,
			a.suspension_remarks AS authentication_suspension_remarks,
			s.current_semester_package,
			s.total_study_plan,
			s.has_paid,
			agl.id AS academic_guidance_lecturer_id,
			agl.name AS academic_guidance_lecturer_name,
			agl.front_title AS academic_guidance_lecturer_front_title,
			agl.back_degree AS academic_guidance_lecturer_back_degree,
			agse.id AS academic_guidance_semester_id,
			agse.semester_start_year AS academic_guidance_semester_start_year,
			s.blood_type,
			s.profile_photo_path,
			s.profile_photo_path_type,
			bp.id AS birth_provice_id,
			bp.name AS birth_provice_name,
			s.height,
			s.weight,
			s.is_color_blind,
			s.use_glasses,
			s.has_complete_teeth,
			s.is_kps_recipient,
			s.work_address,
			s.assurance_number,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			sl.name AS study_level_name,
			s.parent_religion,
			s.parent_nationality,
			s.father_work_address,
			pp.id AS parent_province_id,
			pp.name AS parent_province_name,
			gp.id AS guardian_province_id,
			gp.name AS guardian_province_name,
			s.school_enrollment_year,
			s.school_enrollment_class,
			s.school_type,
			scp.id AS school_province_id,
			scp.name AS school_province_name,
			s.school_status,
			s.school_mathematics_final_exam_score,
			s.school_indonesian_final_exam_score,
			s.school_english_final_exam_score,
			s.school_mathematics_report_score,
			s.school_indonesian_report_score,
			s.school_english_report_score,
			s.transcript_is_archived,
			s.graduation_date,
			s.diploma_number,
			grp.predicate AS graduation_predicate,
			t.title AS thesis_title,
			t.english_title AS thesis_english_title
		FROM students s
		LEFT JOIN regencies br ON br.id = s.birth_regency_id
		LEFT JOIN provinces bp ON bp.id = br.province_id
		LEFT JOIN villages v ON v.id = s.village_id
		LEFT JOIN districts d ON d.id = v.district_id
		LEFT JOIN regencies r ON r.id = d.regency_id
		LEFT JOIN provinces p ON p.id = r.province_id
		LEFT JOIN countries c ON c.id = p.country_id
		LEFT JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		LEFT JOIN curriculums cu ON cu.id = s.curriculum_id
		LEFT JOIN regencies gr ON gr.id = s.guardian_regency_id
		LEFT JOIN provinces gp ON gp.id = gr.province_id
		LEFT JOIN regencies pr ON pr.id = s.parent_regency_id
		LEFT JOIN provinces pp ON pp.id = pr.province_id
		LEFT JOIN authentications a ON a.student_id = s.id
		LEFT JOIN academic_guidance_students ags ON ags.student_id = s.id
		LEFT JOIN semesters agse ON agse.id = ags.semester_id AND agse.is_active IS true
		LEFT JOIN academic_guidances ag ON ag.id = ags.academic_guidance_id
		LEFT JOIN lecturers agl ON agl.id = ag.lecturer_id
		LEFT JOIN provinces scp ON scp.id = s.school_province_id
		LEFT JOIN graduation_predicates grp ON grp.id = s.graduation_predicate_id
		LEFT JOIN theses t ON t.student_id = s.id AND t.status::text = $1
		%s
		WHERE s.id = $2
	`

	getDetailByIdsQuery = `
		SELECT 
			s.id,
			s.name,
			s.sex,
			s.marital_status,
			br.id AS birth_regency_id,
			br.name AS birth_regency_name,
			s.birth_date,
			s.religion,
			s.address,
			s.rt,
			s.rw,
			v.id AS village_id,
			v.name AS village_name,
			d.id AS district_id,
			d.name AS district_name,
			r.id AS regency_id,
			r.name AS regency_name,
			p.id AS province_id,
			p.name AS province_name,
			c.id AS country_id,
			c.name AS country_name,
			s.postal_code,
			s.previous_address,
			s.id_number,
			s.npwp_number,
			s.nisn_number,
			s.residence_type,
			s.transportation_mean,
			s.kps_receiver,
			s.phone_number,
			s.mobile_phone_number,
			s.email,
			s.homepage,
			s.work_type,
			s.work_place,
			s.nationality,
			s.askes_number,
			s.total_brother,
			s.total_sister,
			s.hobby,
			s.experience,
			s.total_dependent,
			s.nim_number,
			s.student_force,
			s.admittance_semester,
			sp.id AS study_program_id,
			sp.name AS study_program_name,
			sp.minimum_graduation_credit AS study_program_minimum_graduation_credit,
			sp.minimum_thesis_credit AS study_program_minimum_thesis_credit,
			dsp.code AS dikti_study_program_code,
			cu.id AS curriculum_id,
			cu.name AS curriculum_name,
			s.admittance_test_number,
			s.admittance_date,
			s.admittance_status,
			s.total_transfer_credit,
			s.previous_college,
			s.previous_study_program,
			s.previous_nim_number,
			s.previous_nim_admittance_year,
			s.status,
			s.is_foreign_student,
			s.college_entrance_type,
			s.class_time,
			s.fund_source,
			s.is_scholarship_grantee,
			s.has_complete_requirement,
			s.school_certificate_type,
			s.school_graduation_year,
			s.school_name,
			s.school_accreditation,
			s.school_address,
			s.school_major,
			s.school_certificate_number,
			s.school_certificate_date,
			s.total_school_final_exam_subject,
			s.school_final_exam_score,
			s.guardian_name,
			s.guardian_birth_date,
			s.guardian_death_date,
			s.guardian_address,
			gr.id AS guardian_regency_id,
			gr.name AS guardian_regency_name,
			s.guardian_postal_code,
			s.guardian_phone_number,
			s.guardian_email,
			s.guardian_final_academic_background,
			s.guardian_occupation,
			s.father_id_number,
			s.father_name,
			s.father_birth_date,
			s.father_death_date,
			s.mother_id_number,
			s.mother_name,
			s.mother_birth_date,
			s.mother_death_date,
			s.parent_address,
			pr.id AS parent_regency_id,
			pr.name AS parent_regency_name,
			s.parent_postal_code,
			s.parent_phone_number,
			s.parent_email,
			s.father_final_academic_background,
			s.father_occupation,
			s.mother_final_academic_background,
			s.mother_occupation,
			s.parent_income,
			s.is_financially_capable,
			a.id AS authentication_id,
			a.is_active AS authentication_is_active,
			a.suspension_remarks AS authentication_suspension_remarks,
			s.current_semester_package,
			s.total_study_plan,
			s.has_paid,
			agl.id AS academic_guidance_lecturer_id,
			agl.name AS academic_guidance_lecturer_name,
			agl.front_title AS academic_guidance_lecturer_front_title,
			agl.back_degree AS academic_guidance_lecturer_back_degree,
			agse.id AS academic_guidance_semester_id,
			agse.semester_start_year AS academic_guidance_semester_start_year,
			s.blood_type,
			s.profile_photo_path,
			s.profile_photo_path_type,
			bp.id AS birth_provice_id,
			bp.name AS birth_provice_name,
			s.height,
			s.weight,
			s.is_color_blind,
			s.use_glasses,
			s.has_complete_teeth,
			s.is_kps_recipient,
			s.work_address,
			s.assurance_number,
			dsp.type AS dikti_study_program_type,
			sl.short_name AS study_level_short_name,
			sl.name AS study_level_name,
			s.parent_religion,
			s.parent_nationality,
			s.father_work_address,
			pp.id AS parent_province_id,
			pp.name AS parent_province_name,
			gp.id AS guardian_province_id,
			gp.name AS guardian_province_name,
			s.school_enrollment_year,
			s.school_enrollment_class,
			s.school_type,
			scp.id AS school_province_id,
			scp.name AS school_province_name,
			s.school_status,
			s.school_mathematics_final_exam_score,
			s.school_indonesian_final_exam_score,
			s.school_english_final_exam_score,
			s.school_mathematics_report_score,
			s.school_indonesian_report_score,
			s.school_english_report_score,
			s.transcript_is_archived,
			s.graduation_date,
			s.diploma_number,
			grp.predicate AS graduation_predicate,
			t.title AS thesis_title,
			t.english_title AS thesis_english_title
		FROM students s
		LEFT JOIN regencies br ON br.id = s.birth_regency_id
		LEFT JOIN provinces bp ON bp.id = br.province_id
		LEFT JOIN villages v ON v.id = s.village_id
		LEFT JOIN districts d ON d.id = v.district_id
		LEFT JOIN regencies r ON r.id = d.regency_id
		LEFT JOIN provinces p ON p.id = r.province_id
		LEFT JOIN countries c ON c.id = p.country_id
		LEFT JOIN study_programs sp ON sp.id = s.study_program_id
		LEFT JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		LEFT JOIN study_levels sl ON sl.id = dsp.study_level_id
		LEFT JOIN curriculums cu ON cu.id = s.curriculum_id
		LEFT JOIN regencies gr ON gr.id = s.guardian_regency_id
		LEFT JOIN provinces gp ON gp.id = gr.province_id
		LEFT JOIN regencies pr ON pr.id = s.parent_regency_id
		LEFT JOIN provinces pp ON pp.id = pr.province_id
		LEFT JOIN authentications a ON a.student_id = s.id
		LEFT JOIN academic_guidance_students ags ON ags.student_id = s.id
		LEFT JOIN semesters agse ON agse.id = ags.semester_id AND agse.is_active IS true
		LEFT JOIN academic_guidances ag ON ag.id = ags.academic_guidance_id
		LEFT JOIN lecturers agl ON agl.id = ag.lecturer_id
		LEFT JOIN provinces scp ON scp.id = s.school_province_id
		LEFT JOIN graduation_predicates grp ON grp.id = s.graduation_predicate_id
		LEFT JOIN theses t ON t.student_id = s.id AND t.status::text = $1
		WHERE s.id IN (SELECT UNNEST($2::uuid[]))
	`

	createQuery = `
		INSERT INTO students (
			name,
			sex,
			marital_status,
			birth_regency_id,
			birth_date,
			religion,
			address,
			rt,
			rw,
			village_id,
			postal_code,
			previous_address,
			id_number,
			npwp_number,
			nisn_number,
			residence_type,
			transportation_mean,
			kps_receiver,
			phone_number,
			mobile_phone_number,
			email,
			homepage,
			work_type,
			work_place,
			nationality,
			askes_number,
			total_brother,
			total_sister,
			hobby,
			experience,
			total_dependent,
			nim_number,
			student_force,
			admittance_semester,
			study_program_id,
			curriculum_id,
			admittance_test_number,
			admittance_date,
			admittance_status,
			total_transfer_credit,
			previous_college,
			previous_study_program,
			previous_nim_number,
			previous_nim_admittance_year,
			status,
			is_foreign_student,
			college_entrance_type,
			class_time,
			fund_source,
			is_scholarship_grantee,
			has_complete_requirement,
			school_certificate_type,
			school_graduation_year,
			school_name,
			school_accreditation,
			school_address,
			school_major,
			school_certificate_number,
			school_certificate_date,
			total_school_final_exam_subject,
			school_final_exam_score,
			guardian_name,
			guardian_birth_date,
			guardian_death_date,
			guardian_address,
			guardian_regency_id,
			guardian_postal_code,
			guardian_phone_number,
			guardian_email,
			guardian_final_academic_background,
			guardian_occupation,
			father_id_number,
			father_name,
			father_birth_date,
			father_death_date,
			mother_id_number,
			mother_name,
			mother_birth_date,
			mother_death_date,
			parent_address,
			parent_regency_id,
			parent_postal_code,
			parent_phone_number,
			parent_email,
			father_final_academic_background,
			father_occupation,
			mother_final_academic_background,
			mother_occupation,
			parent_income,
			is_financially_capable
		) VALUES (
			:name,
			:sex,
			:marital_status,
			:birth_regency_id,
			:birth_date,
			:religion,
			:address,
			:rt,
			:rw,
			:village_id,
			:postal_code,
			:previous_address,
			:id_number,
			:npwp_number,
			:nisn_number,
			:residence_type,
			:transportation_mean,
			:kps_receiver,
			:phone_number,
			:mobile_phone_number,
			:email,
			:homepage,
			:work_type,
			:work_place,
			:nationality,
			:askes_number,
			:total_brother,
			:total_sister,
			:hobby,
			:experience,
			:total_dependent,
			:nim_number,
			:student_force,
			:admittance_semester,
			:study_program_id,
			:curriculum_id,
			:admittance_test_number,
			:admittance_date,
			:admittance_status,
			:total_transfer_credit,
			:previous_college,
			:previous_study_program,
			:previous_nim_number,
			:previous_nim_admittance_year,
			:status,
			:is_foreign_student,
			:college_entrance_type,
			:class_time,
			:fund_source,
			:is_scholarship_grantee,
			:has_complete_requirement,
			:school_certificate_type,
			:school_graduation_year,
			:school_name,
			:school_accreditation,
			:school_address,
			:school_major,
			:school_certificate_number,
			:school_certificate_date,
			:total_school_final_exam_subject,
			:school_final_exam_score,
			:guardian_name,
			:guardian_birth_date,
			:guardian_death_date,
			:guardian_address,
			:guardian_regency_id,
			:guardian_postal_code,
			:guardian_phone_number,
			:guardian_email,
			:guardian_final_academic_background,
			:guardian_occupation,
			:father_id_number,
			:father_name,
			:father_birth_date,
			:father_death_date,
			:mother_id_number,
			:mother_name,
			:mother_birth_date,
			:mother_death_date,
			:parent_address,
			:parent_regency_id,
			:parent_postal_code,
			:parent_phone_number,
			:parent_email,
			:father_final_academic_background,
			:father_occupation,
			:mother_final_academic_background,
			:mother_occupation,
			:parent_income,
			:is_financially_capable
		)
	`

	updateQuery = `
		UPDATE students SET
			name = :name,
			sex = :sex,
			marital_status = :marital_status,
			birth_regency_id = :birth_regency_id,
			birth_date = :birth_date,
			religion = :religion,
			address = :address,
			rt = :rt,
			rw = :rw,
			village_id = :village_id,
			postal_code = :postal_code,
			previous_address = :previous_address,
			id_number = :id_number,
			npwp_number = :npwp_number,
			nisn_number = :nisn_number,
			residence_type = :residence_type,
			transportation_mean = :transportation_mean,
			kps_receiver = :kps_receiver,
			phone_number = :phone_number,
			mobile_phone_number = :mobile_phone_number,
			email = :email,
			homepage = :homepage,
			work_type = :work_type,
			work_place = :work_place,
			nationality = :nationality,
			askes_number = :askes_number,
			total_brother = :total_brother,
			total_sister = :total_sister,
			hobby = :hobby,
			experience = :experience,
			total_dependent = :total_dependent,
			nim_number = :nim_number,
			student_force = :student_force,
			admittance_semester = :admittance_semester,
			study_program_id = :study_program_id,
			curriculum_id = :curriculum_id,
			admittance_test_number = :admittance_test_number,
			admittance_date = :admittance_date,
			admittance_status = :admittance_status,
			total_transfer_credit = :total_transfer_credit,
			previous_college = :previous_college,
			previous_study_program = :previous_study_program,
			previous_nim_number = :previous_nim_number,
			previous_nim_admittance_year = :previous_nim_admittance_year,
			status = :status,
			is_foreign_student = :is_foreign_student,
			college_entrance_type = :college_entrance_type,
			class_time = :class_time,
			fund_source = :fund_source,
			is_scholarship_grantee = :is_scholarship_grantee,
			has_complete_requirement = :has_complete_requirement,
			school_certificate_type = :school_certificate_type,
			school_graduation_year = :school_graduation_year,
			school_name = :school_name,
			school_accreditation = :school_accreditation,
			school_address = :school_address,
			school_major = :school_major,
			school_certificate_number = :school_certificate_number,
			school_certificate_date = :school_certificate_date,
			total_school_final_exam_subject = :total_school_final_exam_subject,
			school_final_exam_score = :school_final_exam_score,
			guardian_name = :guardian_name,
			guardian_birth_date = :guardian_birth_date,
			guardian_death_date = :guardian_death_date,
			guardian_address = :guardian_address,
			guardian_regency_id = :guardian_regency_id,
			guardian_postal_code = :guardian_postal_code,
			guardian_phone_number = :guardian_phone_number,
			guardian_email = :guardian_email,
			guardian_final_academic_background = :guardian_final_academic_background,
			guardian_occupation = :guardian_occupation,
			father_id_number = :father_id_number,
			father_name = :father_name,
			father_birth_date = :father_birth_date,
			father_death_date = :father_death_date,
			mother_id_number = :mother_id_number,
			mother_name = :mother_name,
			mother_birth_date = :mother_birth_date,
			mother_death_date = :mother_death_date,
			parent_address = :parent_address,
			parent_regency_id = :parent_regency_id,
			parent_postal_code = :parent_postal_code,
			parent_phone_number = :parent_phone_number,
			parent_email = :parent_email,
			father_final_academic_background = :father_final_academic_background,
			father_occupation = :father_occupation,
			mother_final_academic_background = :mother_final_academic_background,
			mother_occupation = :mother_occupation,
			parent_income = :parent_income,
			is_financially_capable = :is_financially_capable
		WHERE id = :id
	`

	deleteQuery = `
		DELETE FROM students WHERE id = $1
	`

	getActiveQuery = `
		SELECT 
			s.id,
			s.nim_number,
			s.name,
			s.study_program_id,
			s.curriculum_id,
			s.current_semester_package
			%s
		FROM students s
		%s
		WHERE s.status IN ($1, $2) AND s.study_program_id IS NOT NULL AND s.curriculum_id IS NOT NULL
	`

	updateActiveSemesterPackageQuery = `
		UPDATE students SET current_semester_package = current_semester_package + 1
		WHERE status = $1 AND study_program_id IS NOT NULL AND curriculum_id IS NOT NULL
	`

	bulkUpdateStatusQuery = `
		UPDATE students SET
			status = $1,
			status_reference_number = $2,
			status_date = $3,
			status_purpose = $4,
			status_remarks = $5
		WHERE id IN (SELECT UNNEST($6::uuid[]))
	`

	getStatusSummaryQuery = `
		WITH d AS (
			SELECT UNNEST($1::text[]) AS status
		)
		SELECT 
			ssl.study_program_id, 
			d.status, 
			COUNT(ssl.id) AS total
		FROM d
		LEFT JOIN student_status_logs ssl ON ssl.status = d.status AND ssl.study_program_id IN (SELECT UNNEST($2::uuid[])) AND ssl.semester_id = $3
		LEFT JOIN students s ON s.id = ssl.student_id
		GROUP BY 
			ssl.study_program_id,
			d.status
		ORDER BY d.status
	`

	updateProfileQuery = `
		UPDATE students SET
			profile_photo_path = :profile_photo_path,
			profile_photo_path_type = :profile_photo_path_type,
			sex = :sex,
			birth_regency_id = :birth_regency_id,
			blood_type = :blood_type,
			height = :height,
			weight = :weight,
			is_color_blind = :is_color_blind,
			use_glasses = :use_glasses,
			has_complete_teeth = :has_complete_teeth,
			id_number = :id_number,
			npwp_number = :npwp_number,
			nisn_number = :nisn_number,
			religion = :religion,
			marital_status = :marital_status,
			nationality = :nationality,
			village_id = :village_id,
			rt = :rt,
			rw = :rw,
			postal_code = :postal_code,
			address = :address,
			phone_number = :phone_number,
			mobile_phone_number = :mobile_phone_number,
			email = :email,
			transportation_mean = :transportation_mean,
			is_kps_recipient = :is_kps_recipient,
			fund_source = :fund_source,
			is_scholarship_grantee = :is_scholarship_grantee,
			total_brother = :total_brother,
			total_sister = :total_sister,
			work_type = :work_type,
			work_place = :work_place,
			work_address = :work_address,
			assurance_number = :assurance_number,
			hobby = :hobby
		WHERE id = :id
	`

	updateParentProfileQuery = `
		UPDATE students SET
			father_id_number = :father_id_number,
			father_name = :father_name,
			father_birth_date = :father_birth_date,
			father_death_date = :father_death_date,
			father_final_academic_background = :father_final_academic_background,
			father_occupation = :father_occupation,
			mother_id_number = :mother_id_number,
			mother_name = :mother_name,
			mother_birth_date = :mother_birth_date,
			mother_death_date = :mother_death_date,
			mother_final_academic_background = :mother_final_academic_background,
			mother_occupation = :mother_occupation,
			parent_religion = :parent_religion,
			parent_nationality = :parent_nationality,
			parent_address = :parent_address,
			father_work_address = :father_work_address,
			parent_regency_id = :parent_regency_id,
			parent_postal_code = :parent_postal_code,
			parent_phone_number = :parent_phone_number,
			parent_email = :parent_email,
			is_financially_capable = :is_financially_capable,
			parent_income = :parent_income,
			total_dependent = :total_dependent,
			guardian_name = :guardian_name,
			guardian_birth_date = :guardian_birth_date,
			guardian_death_date = :guardian_death_date,
			guardian_address = :guardian_address,
			guardian_regency_id = :guardian_regency_id,
			guardian_postal_code = :guardian_postal_code,
			guardian_phone_number = :guardian_phone_number,
			guardian_email = :guardian_email,
			guardian_final_academic_background = :guardian_final_academic_background,
			guardian_occupation = :guardian_occupation
		WHERE id = :id
	`

	getPreHighSchoolHistoryByStudentIdsQuery = `
		SELECT
			sphsh.id,
			sphsh.student_id,
			sphsh.level,
			sphsh.name,
			sphsh.graduation_year
		FROM student_pre_high_school_histories sphsh
		WHERE sphsh.student_id IN (SELECT UNNEST($1::uuid[]))
		ORDER BY sphsh.student_id ASC, sphsh.graduation_year DESC
	`

	updateSchoolProfileQuery = `
		UPDATE students SET
			school_enrollment_year = :school_enrollment_year,
			school_graduation_year = :school_graduation_year,
			school_enrollment_class = :school_enrollment_class,
			school_major = :school_major,
			school_type = :school_type,
			school_name = :school_name,
			school_province_id = :school_province_id,
			school_address = :school_address,
			school_certificate_number = :school_certificate_number,
			school_certificate_date = :school_certificate_date,
			school_status = :school_status,
			school_accreditation = :school_accreditation,
			school_final_exam_score = :school_final_exam_score,
			school_mathematics_final_exam_score = :school_mathematics_final_exam_score,
			school_indonesian_final_exam_score = :school_indonesian_final_exam_score,
			school_english_final_exam_score = :school_english_final_exam_score,
			school_mathematics_report_score = :school_mathematics_report_score,
			school_indonesian_report_score = :school_indonesian_report_score,
			school_english_report_score = :school_english_report_score
		WHERE id = :id
	`

	upsertPreHighSchoolHistoryQuery = `
		INSERT INTO student_pre_high_school_histories (
			student_id,
			level,
			name,
			graduation_year
		) VALUES (
			:student_id,
			:level,
			:name,
			:graduation_year
		) ON CONFLICT (student_id, level) DO UPDATE SET
			name = EXCLUDED.name,
			graduation_year = EXCLUDED.graduation_year
	`

	deletePreHighSchoolHistoryExcludingLevelQuery = `
		DELETE FROM student_pre_high_school_histories
		WHERE student_id = $1 %s
	`

	getStudentSubjectQuery = `
		SELECT
			s.id AS subject_id,
			s.code AS subject_code,
			s.name AS subject_name,
			s.english_name AS subject_english_name,
			s.theory_credit AS subject_theory_credit,
			s.practicum_credit AS subject_practicum_credit,
			s.field_practicum_credit AS subject_field_practicum_credit,
			gs.id AS grade_semester_id,
			gs.semester_start_year AS grade_semester_start_year,
			gs.semester_type AS grade_semester_type,
			ss.grade_point,
			ss.grade_code,
			s.is_mandatory AS subject_is_mandatory,
			ss.semester_package,
			(s.theory_credit + s.practicum_credit + s.field_practicum_credit) AS subject_total_credit,
			s.type AS subject_type
		FROM student_subjects ss
		JOIN subjects s ON s.id = ss.subject_id
		JOIN semesters gs ON gs.id = ss.grade_semester_id
	`

	countStudentSubjectQuery = `
		SELECT COUNT(1)
		FROM student_subjects ss
		JOIN subjects s ON s.id = ss.subject_id
		JOIN semesters gs ON gs.id = ss.grade_semester_id
	`

	updatePaymentQuery = `
		UPDATE students SET has_paid = true, paid_by = $1
		WHERE id IN (SELECT UNNEST($2::uuid[]))
	`

	getPaymentLogQuery = `
		SELECT
			se.id AS semester_id,
			se.semester_type,
			se.semester_start_year,
			spl.created_at
		FROM student_payment_logs spl
		JOIN semesters se ON se.id = spl.semester_id
		WHERE spl.student_id = $1
	`

	bulkCreateQuery = `
		INSERT INTO students (
			id,
			nim_number,
			name,
			sex,
			marital_status,
			birth_regency_id,
			birth_date,
			religion,
			address,
			rt,
			rw,
			village_id,
			postal_code,
			id_number,
			nisn_number,
			mobile_phone_number,
			nationality,
			study_program_id,
			school_name,
			school_address,
			school_province_id,
			school_major,
			school_type,
			school_graduation_year,
			father_name,
			father_id_number,
			father_birth_date,
			father_final_academic_background,
			father_occupation,
			mother_name,
			mother_id_number,
			mother_birth_date,
			mother_final_academic_background,
			mother_occupation,
			guardian_name,
			guardian_id_number,
			guardian_birth_date,
			guardian_final_academic_background,
			guardian_occupation
		) VALUES (
			:id,
			:nim_number,
			:name,
			:sex,
			:marital_status,
			:birth_regency_id,
			:birth_date,
			:religion,
			:address,
			:rt,
			:rw,
			:village_id,
			:postal_code,
			:id_number,
			:nisn_number,
			:mobile_phone_number,
			:nationality,
			:study_program_id,
			:school_name,
			:school_address,
			:school_province_id,
			:school_major,
			:school_type,
			:school_graduation_year,
			:father_name,
			:father_id_number,
			:father_birth_date,
			:father_final_academic_background,
			:father_occupation,
			:mother_name,
			:mother_id_number,
			:mother_birth_date,
			:mother_final_academic_background,
			:mother_occupation,
			:guardian_name,
			:guardian_id_number,
			:guardian_birth_date,
			:guardian_final_academic_background,
			:guardian_occupation
		) ON CONFLICT (nim_number) DO NOTHING;
	`

	convertGradeQuery = `
		UPDATE student_subjects SET
			grade_semester_id = $3,
			grade_point = $4,
			grade_code = $5,
			mbkm_subject_id = $6
		WHERE student_id = $1 AND subject_id IN (SELECT UNNEST($2::uuid[]))
	`
)
