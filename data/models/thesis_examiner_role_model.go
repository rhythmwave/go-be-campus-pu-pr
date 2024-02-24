package models

type GetThesisExaminerRole struct {
	Id      string `db:"id"`
	Name    string `db:"name"`
	Remarks string `db:"remarks"`
	Sort    uint32 `db:"sort"`
}

type CreateThesisExaminerRole struct {
	Name      string `db:"name"`
	Remarks   string `db:"remarks"`
	Sort      uint32 `db:"sort"`
	CreatedBy string `db:"created_by"`
}

type UpdateThesisExaminerRole struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	Remarks   string `db:"remarks"`
	Sort      uint32 `db:"sort"`
	UpdatedBy string `db:"updated_by"`
}
