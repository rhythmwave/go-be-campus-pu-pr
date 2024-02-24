package lesson_plan

const (
	getListQuery = `
		SELECT
			lp.id,
			lp.meeting_order,
			lp.lesson,
			lp.english_lesson
		FROM lesson_plans lp
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM lesson_plans lp
	`

	getDetailQuery = `
		SELECT
			lp.id,
			lp.meeting_order,
			lp.lesson,
			lp.english_lesson
		FROM lesson_plans lp
		WHERE lp.id = $1
	`

	createQuery = `
		INSERT INTO lesson_plans (
			subject_id,
			meeting_order,
			lesson,
			english_lesson,
			created_by
		) VALUES ($1, $2, $3, $4, $5);
	`

	updateQuery = `
		UPDATE lesson_plans SET
			meeting_order = $1,
			lesson = $2,
			english_lesson = $3,
			updated_by = $4
		WHERE id = $5
	`

	deleteQuery = `
		DELETE FROM lesson_plans WHERE id = $1
	`
)
