package models

type GetSubjectGradeComponent struct {
	Id         string  `db:"id"`
	SubjectId  string  `db:"subject_id"`
	Name       string  `db:"name"`
	Percentage float64 `db:"percentage"`
	IsActive   bool    `db:"is_active"`
}

type CreateSubjectGradeComponent struct {
	SubjectId  string  `db:"subject_id"`
	Name       string  `db:"name"`
	Percentage float64 `db:"percentage"`
	IsActive   bool    `db:"is_active"`
	CreatedBy  string  `db:"created_by"`
}
