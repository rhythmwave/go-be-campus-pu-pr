package shared_file

const (
	getListQuery = `
		SELECT
			sf.id,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			sf.title,
			sf.file_path,
			sf.file_path_type,
			sf.remarks,
			sf.is_approved,
			sf.created_at
		FROM shared_files sf
		JOIN lecturers l ON l.id = sf.lecturer_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM shared_files sf
	`

	getDetailQuery = `
		SELECT 
			sf.id,
			l.id AS lecturer_id,
			l.name AS lecturer_name,
			l.front_title AS lecturer_front_title,
			l.back_degree AS lecturer_back_degree,
			sf.title,
			sf.file_path,
			sf.file_path_type,
			sf.remarks,
			sf.is_approved,
			sf.created_at
		FROM shared_files sf
		JOIN lecturers l ON l.id = sf.lecturer_id
		WHERE sf.id = $1
	`

	createQuery = `
		INSERT INTO shared_files (
			lecturer_id,
			title,
			file_path,
			file_path_type,
			remarks
		) VALUES (
			:lecturer_id,
			:title,
			:file_path,
			:file_path_type,
			:remarks
		);
	`

	updateQuery = `
		UPDATE shared_files SET
			title = :title,
			remarks = :remarks
		WHERE id = :id
	`

	approveQuery = `
		UPDATE shared_files SET is_approved = true WHERE id = $1
	`

	deleteQuery = `
		DELETE FROM shared_files WHERE id = $1
	`
)
