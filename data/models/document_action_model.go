package models

type GetDocumentAction struct {
	Id            string `db:"id"`
	Action        string `db:"action"`
	EnglishAction string `db:"english_action"`
}

type CreateDocumentAction struct {
	Action        string `db:"action"`
	EnglishAction string `db:"english_action"`
	CreatedBy     string `db:"created_by"`
}

type UpdateDocumentAction struct {
	Id            string `db:"id"`
	Action        string `db:"action"`
	EnglishAction string `db:"english_action"`
	UpdatedBy     string `db:"updated_by"`
}
