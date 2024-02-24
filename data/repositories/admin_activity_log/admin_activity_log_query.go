package admin_activity_log

const (
	getListQuery = `
		SELECT
			aal.id,
			aal.admin_id,
			aal.admin_name,
			aal.admin_username,
			aal.module,
			aal.action,
			aal.ip_address,
			aal.user_agent,
			aal.execution_time,
			aal.memory_usage,
			aal.created_at
		FROM admin_activity_logs aal
	`

	countListQuery = `
		SELECT COUNT(1)
		FROM admin_activity_logs aal
	`

	createQuery = `
		INSERT INTO admin_activity_logs (
			admin_id,
			admin_name,
			admin_username,
			module,
			action,
			ip_address,
			user_agent,
			execution_time,
			memory_usage
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
)
