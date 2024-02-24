package models

type GetExpertiseGroup struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type CreateExpertiseGroup struct {
	Name      string `db:"name"`
	CreatedBy string `db:"created_by"`
}

type UpdateExpertiseGroup struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	UpdatedBy string `db:"updated_by"`
}
