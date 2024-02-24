package models

type GetThesisSupervisorRole struct {
	Id   string `db:"id"`
	Name string `db:"name"`
	Sort uint32 `db:"sort"`
}

type CreateThesisSupervisorRole struct {
	Name      string `db:"name"`
	Sort      uint32 `db:"sort"`
	CreatedBy string `db:"created_by"`
}

type UpdateThesisSupervisorRole struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	Sort      uint32 `db:"sort"`
	UpdatedBy string `db:"updated_by"`
}
