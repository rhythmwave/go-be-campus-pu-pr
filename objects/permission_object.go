package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetPermission struct {
	Id   string
	Name string
}

type PermissionListWithPagination struct {
	Pagination common.Pagination
	Data       []GetPermission
}
