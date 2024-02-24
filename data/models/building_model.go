package models

import (
	"database/sql"
)

type GetBuilding struct {
	Id               string  `db:"id"`
	FacultyId        *string `db:"faculty_id"`
	FacultyName      *string `db:"faculty_name"`
	MajorFacultyId   *string `db:"major_faculty_id"`
	MajorFacultyName *string `db:"major_faculty_name"`
	MajorId          *string `db:"major_id"`
	MajorName        *string `db:"major_name"`
	Code             string  `db:"code"`
	Name             string  `db:"name"`
}

type CreateBuilding struct {
	FacultyId sql.NullString `db:"faculty_id"`
	MajorId   sql.NullString `db:"major_id"`
	Code      string         `db:"code"`
	Name      string         `db:"name"`
	CreatedBy string         `db:"created_by"`
}

type UpdateBuilding struct {
	Id        string         `db:"id"`
	FacultyId sql.NullString `db:"faculty_id"`
	MajorId   sql.NullString `db:"major_id"`
	Code      string         `db:"code"`
	Name      string         `db:"name"`
	UpdatedBy string         `db:"updated_by"`
}
