package learning_achievement_category

const (
	getListQuery = `
		SELECT
			lac.id,
			c.id AS curriculum_id,
			c.name AS curriculum_name,
			lac.name,
			lac.english_name
		FROM learning_achievement_categories lac
		JOIN curriculums c ON c.id = lac.curriculum_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM learning_achievement_categories lac
		JOIN curriculums c ON c.id = lac.curriculum_id
	`

	getDetailQuery = `
		SELECT 
			lac.id,
			c.id AS curriculum_id,
			c.name AS curriculum_name,
			lac.name,
			lac.english_name
		FROM learning_achievement_categories lac
		JOIN curriculums c ON c.id = lac.curriculum_id
		WHERE lac.id = $1
	`

	createQuery = `
		INSERT INTO learning_achievement_categories (
			curriculum_id,
			name,
			english_name,
			created_by
		) VALUES ($1, $2, $3, $4);
	`

	updateQuery = `
		UPDATE learning_achievement_categories SET
			name = $1,
			english_name = $2,
			updated_by = $3
		WHERE id = $4
	`

	deleteQuery = `
		DELETE FROM learning_achievement_categories WHERE id = $1
	`
)
