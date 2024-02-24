package models

type GetPermissionList struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type GetPermissionByRoleIds struct {
	GetPermissionList
	RoleId string `db:"role_id"`
}
