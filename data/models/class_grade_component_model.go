package models

type GetClassGradeComponent struct {
	Id         string  `db:"id"`
	ClassId    string  `db:"class_id"`
	Name       string  `db:"name"`
	Percentage float64 `db:"percentage"`
	IsActive   bool    `db:"is_active"`
}

type CreateClassGradeComponent struct {
	ClassId    string  `db:"class_id"`
	Name       string  `db:"name"`
	Percentage float64 `db:"percentage"`
	IsActive   bool    `db:"is_active"`
}
