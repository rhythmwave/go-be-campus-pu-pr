package learning_achievement

const (
	getListQuery = `
		SELECT
			la.id,
			lac.id AS learning_achievement_category_id,
			lac.name AS learning_achievement_category_name,
			lac.english_name AS learning_achievement_category_english_name,
			la.name,
			la.english_name
		FROM learning_achievements la
		JOIN learning_achievement_categories lac ON lac.id = la.learning_achievement_category_id
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM learning_achievements la
		JOIN learning_achievement_categories lac ON lac.id = la.learning_achievement_category_id
	`

	getDetailQuery = `
		SELECT 
			la.id,
			lac.id AS learning_achievement_category_id,
			lac.name AS learning_achievement_category_name,
			lac.english_name AS learning_achievement_category_english_name,
			la.name,
			la.english_name
		FROM learning_achievements la
		JOIN learning_achievement_categories lac ON lac.id = la.learning_achievement_category_id
		WHERE la.id = $1
	`

	createQuery = `
		INSERT INTO learning_achievements (
			learning_achievement_category_id,
			name,
			english_name,
			created_by
		) VALUES ($1, $2, $3, $4);
	`

	updateQuery = `
		UPDATE learning_achievements SET
			name = $1,
			english_name = $2,
			updated_by = $3
		WHERE id = $4
	`

	deleteQuery = `
		DELETE FROM learning_achievements WHERE id = $1
	`
)
