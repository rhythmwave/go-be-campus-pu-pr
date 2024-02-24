package permission

const (
	getListQuery = `
		SELECT p.id, p.name FROM permissions p
	`

	countListQuery = `
		SELECT COUNT(1) FROM permissions p
	`

	getByRoleIdsQuery = `
		SELECT
			p.id,
			p.name,
			rp.role_id
		FROM permissions p
		JOIN role_permission rp ON rp.permission_id = p.id
		WHERE rp.role_id IN (SELECT UNNEST($1::uuid[]))
		ORDER BY p.name, rp.role_id
	`
)
