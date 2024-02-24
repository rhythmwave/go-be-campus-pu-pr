package models

type GetRoleList struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

type CreateRole struct {
	Name          string `db:"name"`
	Description   string `db:"description"`
	StudyPrograms []string
	Permissions   []string
	CreatedBy     string `db:"created_by"`
}

type UpdateRole struct {
	Id            string `db:"id"`
	Name          string `db:"name"`
	Description   string `db:"description"`
	StudyPrograms []string
	Permissions   []string
	UpdatedBy     string `db:"updated_by"`
}
