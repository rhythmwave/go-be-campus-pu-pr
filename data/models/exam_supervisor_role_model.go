package models

type GetExamSupervisorRole struct {
	Id   string `db:"id"`
	Name string `db:"name"`
	Sort uint32 `db:"sort"`
}

type CreateExamSupervisorRole struct {
	Name      string `db:"name"`
	Sort      uint32 `db:"sort"`
	CreatedBy string `db:"created_by"`
}

type UpdateExamSupervisorRole struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	Sort      uint32 `db:"sort"`
	UpdatedBy string `db:"updated_by"`
}
