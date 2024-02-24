package models

import "database/sql"

type GetAuthentication struct {
	Id                string  `db:"id"`
	Username          string  `db:"username"`
	Password          string  `db:"password"`
	AdminId           *string `db:"admin_id"`
	LecturerId        *string `db:"lecturer_id"`
	StudentId         *string `db:"student_id"`
	AdminRoleId       *string `db:"admin_role_id"`
	AdminRoleName     *string `db:"admin_role_name"`
	AdminName         *string `db:"admin_name"`
	LecturerName      *string `db:"lecturer_name"`
	StudentName       *string `db:"student_name"`
	IsActive          bool    `db:"is_active"`
	SuspensionRemarks *string `db:"suspension_remarks"`
	IdSso             *string `db:"id_sso"`
	SsoRefreshToken   *string `db:"sso_refresh_token"`
}

type CreateAuthentication struct {
	Username   string         `db:"username"`
	Password   string         `db:"password"`
	AdminId    sql.NullString `db:"admin_id"`
	LecturerId sql.NullString `db:"lecturer_id"`
	StudentId  sql.NullString `db:"student_id"`
}
