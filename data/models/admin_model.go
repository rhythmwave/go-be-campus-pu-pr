package models

import "database/sql"

type GetAdmin struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Username string `db:"username"`
	RoleId   string `db:"role_id"`
	RoleName string `db:"role_name"`
}

type CreateAdmin struct {
	Username  string `db:"username"`
	Name      string `db:"name"`
	Password  string `db:"password"`
	RoleId    string `db:"role_id"`
	CreatedBy string `db:"created_by"`
}

type UpdateAdmin struct {
	Id        string         `db:"id"`
	Username  string         `db:"username"`
	Name      string         `db:"name"`
	Password  sql.NullString `db:"password"`
	RoleId    string         `db:"role_id"`
	UpdatedBy string         `db:"updated_by"`
}
