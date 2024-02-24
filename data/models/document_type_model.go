package models

type GetDocumentType struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

type CreateDocumentType struct {
	Name      string `db:"name"`
	CreatedBy string `db:"created_by"`
}

type UpdateDocumentType struct {
	Id        string `db:"id"`
	Name      string `db:"name"`
	UpdatedBy string `db:"updated_by"`
}
