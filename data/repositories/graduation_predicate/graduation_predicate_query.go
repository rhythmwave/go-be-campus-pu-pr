package graduation_predicate

const (
	getListQuery = `
		SELECT
			gp.id,
			gp.predicate,
			gp.minimum_gpa,
			gp.maximum_study_semester,
			gp.repeat_course_limit,
			gp.below_minimum_grade_point_limit
		FROM graduation_predicates gp
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM graduation_predicates gp
	`

	getDetailQuery = `
		SELECT 
			gp.id,
			gp.predicate,
			gp.minimum_gpa,
			gp.maximum_study_semester,
			gp.repeat_course_limit,
			gp.below_minimum_grade_point_limit
		FROM graduation_predicates gp
		WHERE gp.id = $1
	`

	createQuery = `
		INSERT INTO graduation_predicates (
			predicate,
			minimum_gpa,
			maximum_study_semester,
			repeat_course_limit,
			below_minimum_grade_point_limit,
			created_by
		) VALUES ($1, $2, $3, $4, $5, $6);
	`

	updateQuery = `
		UPDATE graduation_predicates SET
			predicate = $1,
			minimum_gpa = $2,
			maximum_study_semester = $3,
			repeat_course_limit = $4,
			below_minimum_grade_point_limit = $5,
			updated_by = $6
		WHERE id = $7
	`

	deleteQuery = `
		DELETE FROM graduation_predicates WHERE id = $1
	`
)
