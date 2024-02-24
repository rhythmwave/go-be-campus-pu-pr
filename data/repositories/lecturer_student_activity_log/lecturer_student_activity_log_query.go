package lecturer_student_activity_log

const (
	getListQuery = `
		SELECT
			lsal.id,
			lsal.student_id,
			lsal.student_name,
			lsal.student_username,
			lsal.lecturer_id,
			lsal.lecturer_name,
			lsal.lecturer_username,
			lsal.module,
			lsal.action,
			lsal.ip_address,
			lsal.user_agent,
			lsal.execution_time,
			lsal.memory_usage,
			lsal.created_at
		FROM lecturer_student_activity_logs lsal
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM lecturer_student_activity_logs lsal
	`

	createQuery = `
		INSERT INTO lecturer_student_activity_logs (
			lecturer_id,
			lecturer_name,
			lecturer_username,
			student_id,
			student_name,
			student_username,
			module,
			action,
			ip_address,
			user_agent,
			execution_time,
			memory_usage
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
)
