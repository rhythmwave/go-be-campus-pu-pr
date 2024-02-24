package models

type CreateSubjectEquivalence struct {
	SubjectId           string `db:"subject_id"`
	EquivalentSubjectId string `db:"equivalent_subject_id"`
	CreatedBy           string `db:"created_by"`
}
