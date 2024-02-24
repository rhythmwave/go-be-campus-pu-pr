package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetAdmin struct {
	Id            string
	Username      string
	Name          string
	RoleId        string
	RoleName      string
	StudyPrograms []GetStudyProgram
	Permissions   []GetPermission
}

type AdminListWithPagination struct {
	Pagination common.Pagination
	Data       []GetAdmin
}

type CreateAdmin struct {
	Username string
	Name     string
	Password string
	RoleId   string
}

type UpdateAdmin struct {
	Id       string
	Username string
	Name     string
	Password string
	RoleId   string
}
