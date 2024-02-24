package subject_category

const (
	getListQuery = `
		SELECT
			sc.id,
			sc.code,
			sc.name
		FROM subject_categories sc
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM subject_categories sc
	`

	getDetailQuery = `
		SELECT 
			sc.id,
			sc.code,
			sc.name
		FROM subject_categories sc
		WHERE sc.id = $1
	`

	createQuery = `
		INSERT INTO subject_categories (
			code,
			name,
			created_by
		) VALUES ($1, $2, $3);
	`

	updateQuery = `
		UPDATE subject_categories SET
			code = $1,
			name = $2,
			updated_by = $3
		WHERE id = $4
	`

	deleteQuery = `
		DELETE FROM subject_categories WHERE id = $1
	`
)
