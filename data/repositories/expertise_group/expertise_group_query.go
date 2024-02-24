package expertise_group

const (
	getListQuery = `
		SELECT
			dt.id,
			dt.name
		FROM expertise_groups dt
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM expertise_groups dt
	`

	getDetailQuery = `
		SELECT 
			dt.id,
			dt.name
		FROM expertise_groups dt
		WHERE dt.id = $1
	`

	createQuery = `
		INSERT INTO expertise_groups (
			name,
			created_by
		) VALUES ($1, $2);
	`

	updateQuery = `
		UPDATE expertise_groups SET
			name = $1,
			updated_by = $2
		WHERE id = $3
	`

	deleteQuery = `
		DELETE FROM expertise_groups WHERE id = $1
	`
)
