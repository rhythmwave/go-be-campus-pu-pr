package announcement

const (
	getListQuery = `
		SELECT 
			a.id,
			a.type,
			a.title,
			a.announcement_date,
			a.file_path,
			a.file_path_type,
			a.file_title,
			a.content,
			a.for_lecturer,
			a.for_student
		FROM announcements a
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM announcements a
	`

	getAnnouncementStudyProgramByAnnouncementIdsQuery = `
		SELECT
			a.id AS announcement_id,
			a.title AS announcement_title,
			sp.id AS study_program_id,
			sp.name AS study_program_name
		FROM announcement_study_programs asp
		JOIN announcements a ON a.id = asp.announcement_id
		JOIN study_programs sp ON sp.id = asp.study_program_id
		WHERE a.id IN (SELECT UNNEST($1::uuid[]))
	`

	getAnnouncementStudyProgramByStudyProgramIdsQuery = `
		SELECT
			a.id AS announcement_id,
			a.title AS announcement_title,
			sp.id AS study_program_id,
			sp.name AS study_program_name
		FROM announcement_study_programs asp
		JOIN announcements a ON a.id = asp.announcement_id
		JOIN study_programs sp ON sp.id = asp.study_program_id
		WHERE sp.id IN (SELECT UNNEST($1::uuid[]))
	`

	getDetailQuery = `
		SELECT 
			a.id,
			a.type,
			a.title,
			a.announcement_date,
			a.file_path,
			a.file_path_type,
			a.file_title,
			a.content,
			a.for_lecturer,
			a.for_student
		FROM announcements a
		WHERE a.id = $1
	`

	createQuery = `
		INSERT INTO announcements (
			id,
			type,
			title,
			announcement_date,
			file_path,
			file_path_type,
			file_title,
			content,
			for_lecturer,
			for_student,
			created_by
		) VALUES (
			:id,
			:type,
			:title,
			:announcement_date,
			:file_path,
			:file_path_type,
			:file_title,
			:content,
			:for_lecturer,
			:for_student,
			:created_by
		)
	`

	updateQuery = `
		UPDATE announcements SET
			type = :type,
			title = :title,
			announcement_date = :announcement_date,
			file_path = :file_path,
			file_path_type = :file_path_type,
			file_title = :file_title,
			content = :content,
			for_lecturer = :for_lecturer,
			for_student = :for_student,
			updated_by = :updated_by
		WHERE id = :id
	`

	deleteAnnouncementStudyProgramExcludingStudyProgramIdsQuery = `
		DELETE FROM announcement_study_programs WHERE announcement_id = $1 %s
	`

	upsertAnnouncementStudyProgramQuery = `
		INSERT INTO announcement_study_programs (
			announcement_id,
			study_program_id
		) VALUES (
			:announcement_id,
			:study_program_id
		) ON CONFLICT (announcement_id, study_program_id) DO NOTHING
	`

	deleteQuery = `
		DELETE FROM announcements WHERE id = $1
	`
)
