package credit_quota

const (
	getListQuery = `
		SELECT
			cq.id,
			cq.minimum_grade_point,
			cq.maximum_grade_point,
			cq.maximum_credit
		FROM credit_quotas cq
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM credit_quotas cq
	`

	getDetailQuery = `
		SELECT 
			cq.id,
			cq.minimum_grade_point,
			cq.maximum_grade_point,
			cq.maximum_credit
		FROM credit_quotas cq
		WHERE cq.id = $1
	`

	createQuery = `
		INSERT INTO credit_quotas (
			minimum_grade_point,
			maximum_grade_point,
			maximum_credit,
			created_by
		) VALUES ($1, $2, $3, $4);
	`

	updateQuery = `
		UPDATE credit_quotas SET
			minimum_grade_point = $1,
			maximum_grade_point = $2,
			maximum_credit = $3,
			updated_by = $4
		WHERE id = $5
	`

	deleteQuery = `
		DELETE FROM credit_quotas WHERE id = $1
	`
)
