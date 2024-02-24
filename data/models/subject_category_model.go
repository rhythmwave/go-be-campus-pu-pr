package models

type GetSubjectCategory struct {
	Id   string `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
}

type CreateSubjectCategory struct {
	Code      string `db:"code"`
	Name      string `db:"name"`
	CreatedBy string `db:"created_by"`
}

type UpdateSubjectCategory struct {
	Id        string `db:"id"`
	Code      string `db:"code"`
	Name      string `db:"name"`
	UpdatedBy string `db:"updated_by"`
}
