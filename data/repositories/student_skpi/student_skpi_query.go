package student_skpi

const (
	getListQuery = `
		SELECT
			ss.id,
			s.id AS student_id,
			s.nim_number AS student_nim_number,
			s.name AS student_name,
			sp.id AS student_study_program_id,
			sp.name AS student_study_program_name,
			dsp.code AS student_dikti_study_program_code,
			ss.is_approved
		FROM student_skpi ss
		JOIN students s ON s.id = ss.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM student_skpi ss
		JOIN students s ON s.id = ss.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
	`

	getByIdQuery = `
		SELECT
			ss.id,
			s.id AS student_id,
			s.nim_number AS student_nim_number,
			s.name AS student_name,
			sp.id AS student_study_program_id,
			sp.name AS student_study_program_name,
			dsp.code AS student_dikti_study_program_code,
			ss.skpi_number,
			ss.is_approved,
			ss.achievement_path,
			ss.achievement_path_type,
			ss.organization_path,
			ss.organization_path_type,
			ss.certificate_path,
			ss.certificate_path_type,
			ss.language_path,
			ss.language_path_type
		FROM student_skpi ss
		JOIN students s ON s.id = ss.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		WHERE ss.id = $1
	`

	getByStudentIdQuery = `
		SELECT
			ss.id,
			s.id AS student_id,
			s.nim_number AS student_nim_number,
			s.name AS student_name,
			sp.id AS student_study_program_id,
			sp.name AS student_study_program_name,
			dsp.code AS student_dikti_study_program_code,
			ss.skpi_number,
			ss.is_approved,
			ss.achievement_path,
			ss.achievement_path_type,
			ss.organization_path,
			ss.organization_path_type,
			ss.certificate_path,
			ss.certificate_path_type,
			ss.language_path,
			ss.language_path_type
		FROM student_skpi ss
		JOIN students s ON s.id = ss.student_id
		JOIN study_programs sp ON sp.id = s.study_program_id
		JOIN dikti_study_programs dsp ON dsp.id = sp.dikti_study_program_id
		WHERE ss.student_id = $1
	`

	getStudentSkpiAchievementByStudentSkpiIdQuery = `
		SELECT
			ssa.id,
			ssa.name,
			ssa.year
		FROM student_skpi_achievements ssa
		WHERE ssa.student_skpi_id = $1
	`

	getStudentSkpiOrganizationByStudentSkpiIdQuery = `
		SELECT
			sso.id,
			sso.name,
			sso.position,
			sso.service_length
		FROM student_skpi_organizations sso
		WHERE sso.student_skpi_id = $1
	`

	getStudentSkpiCertificateByStudentSkpiIdQuery = `
		SELECT
			ssc.id,
			ssc.name
		FROM student_skpi_certificates ssc
		WHERE ssc.student_skpi_id = $1
	`

	getStudentSkpiCharacterBuildingByStudentSkpiIdQuery = `
		SELECT
			sscb.id,
			sscb.name
		FROM student_skpi_character_buildings sscb
		WHERE sscb.student_skpi_id = $1
	`

	getStudentSkpiInternshipByStudentSkpiIdQuery = `
		SELECT
			ssi.id,
			ssi.name
		FROM student_skpi_internships ssi
		WHERE ssi.student_skpi_id = $1
	`

	getStudentSkpiLanguageByStudentSkpiIdQuery = `
		SELECT
			ssl.id,
			ssl.name,
			ssl.score,
			ssl.date
		FROM student_skpi_languages ssl
		WHERE ssl.student_skpi_id = $1
	`

	upsertStudentSkpiQuery = `
		INSERT INTO student_skpi (
			student_id,
			achievement_path,
			achievement_path_type,
			organization_path,
			organization_path_type,
			certificate_path,
			certificate_path_type,
			language_path,
			language_path_type
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (student_id) DO UPDATE SET
			achievement_path = EXCLUDED.achievement_path,
			achievement_path_type = EXCLUDED.achievement_path_type,
			organization_path = EXCLUDED.organization_path,
			organization_path_type = EXCLUDED.organization_path_type,
			certificate_path = EXCLUDED.certificate_path,
			certificate_path_type = EXCLUDED.certificate_path_type,
			language_path = EXCLUDED.language_path,
			language_path_type = EXCLUDED.language_path_type
		RETURNING id
	`

	deleteStudentSkpiAchievementExcludingNameQuery = `
		DELETE FROM student_skpi_achievements
		WHERE student_skpi_id = $1 %s
	`

	deleteStudentSkpiOrganizationExcludingNameQuery = `
		DELETE FROM student_skpi_organizations
		WHERE student_skpi_id = $1 %s
	`

	deleteStudentSkpiCertificateExcludingNameQuery = `
		DELETE FROM student_skpi_certificates
		WHERE student_skpi_id = $1 %s
	`

	deleteStudentSkpiCharacterBuildingExcludingNameQuery = `
		DELETE FROM student_skpi_character_buildings
		WHERE student_skpi_id = $1 %s
	`

	deleteStudentSkpiInternshipExcludingNameQuery = `
		DELETE FROM student_skpi_internships
		WHERE student_skpi_id = $1 %s
	`

	deleteStudentSkpiLanguageExcludingNameQuery = `
		DELETE FROM student_skpi_languages
		WHERE student_skpi_id = $1 %s
	`

	upsertStudentSkpiAchievementQuery = `
		INSERT INTO student_skpi_achievements (
			student_skpi_id,
			name,
			year
		) VALUES (
			:student_skpi_id,
			:name,
			:year
		) ON CONFLICT (student_skpi_id, name) DO UPDATE SET
			year = EXCLUDED.year
	`

	upsertStudentSkpiOrganizationQuery = `
		INSERT INTO student_skpi_organizations (
			student_skpi_id,
			name,
			position,
			service_length
		) VALUES (
			:student_skpi_id,
			:name,
			:position,
			:service_length
		) ON CONFLICT (student_skpi_id, name) DO UPDATE SET
			service_length = EXCLUDED.service_length
	`

	upsertStudentSkpiCertificateQuery = `
		INSERT INTO student_skpi_certificates (
			student_skpi_id,
			name
		) VALUES (
			:student_skpi_id,
			:name
		) ON CONFLICT (student_skpi_id, name) DO NOTHING
	`

	upsertStudentSkpiCharacterBuildingQuery = `
		INSERT INTO student_skpi_character_buildings (
			student_skpi_id,
			name
		) VALUES (
			:student_skpi_id,
			:name
		) ON CONFLICT (student_skpi_id, name) DO NOTHING
	`

	upsertStudentSkpiInternshipQuery = `
		INSERT INTO student_skpi_internships (
			student_skpi_id,
			name
		) VALUES (
			:student_skpi_id,
			:name
		) ON CONFLICT (student_skpi_id, name) DO NOTHING
	`

	upsertStudentSkpiLanguageQuery = `
		INSERT INTO student_skpi_languages (
			student_skpi_id,
			name,
			score,
			date
		) VALUES (
			:student_skpi_id,
			:name,
			:score,
			:date
		) ON CONFLICT (student_skpi_id, name) DO UPDATE SET
			score = EXCLUDED.score,
			date = EXCLUDED.date
	`

	approveQuery = `
		UPDATE student_skpi SET
			skpi_number = $1,
			is_approved = true
		WHERE id = $2
	`

	deleteQuery = `
		DELETE FROM student_skpi WHERE id = $1
	`
)
