package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetRole struct {
	Id            string
	Name          string
	Description   string
	StudyPrograms []GetStudyProgram
	Permissions   []GetPermission
}

type RoleListWithPagination struct {
	Pagination common.Pagination
	Data       []GetRole
}

type CreateRole struct {
	Name            string
	Description     string
	StudyProgramIds []string
	PermissionIds   []string
}

type UpdateRole struct {
	Id              string
	Name            string
	Description     string
	StudyProgramIds []string
	PermissionIds   []string
}
