package models

import (
	"database/sql"
)

type GetOfficer struct {
	Id                 string  `db:"id"`
	IdNationalLecturer *string `db:"id_national_lecturer"`
	Name               string  `db:"name"`
	Title              *string `db:"title"`
	EnglishTitle       *string `db:"english_title"`
	StudyProgramId     *string `db:"study_program_id"`
	StudyProgramName   *string `db:"study_program_name"`
	SignaturePath      *string `db:"signature_path"`
	SignaturePathType  *string `db:"signature_path_type"`
	ShowSignature      bool    `db:"show_signature"`
	EmployeeNo         *string `db:"employee_no"`
}

type CreateOfficer struct {
	IdNationalLecturer sql.NullString `db:"id_national_lecturer"`
	Name               string         `db:"name"`
	Title              sql.NullString `db:"title"`
	EnglishTitle       sql.NullString `db:"english_title"`
	StudyProgramId     sql.NullString `db:"study_program_id"`
	SignaturePath      sql.NullString `db:"signature_path"`
	SignaturePathType  sql.NullString `db:"signature_path_type"`
	ShowSignature      bool           `db:"show_signature"`
	EmployeeNo         sql.NullString `db:"employee_no"`
	CreatedBy          string         `db:"created_by"`
}

type UpdateOfficer struct {
	Id                 string         `db:"id"`
	IdNationalLecturer sql.NullString `db:"id_national_lecturer"`
	Name               string         `db:"name"`
	Title              sql.NullString `db:"title"`
	EnglishTitle       sql.NullString `db:"english_title"`
	StudyProgramId     sql.NullString `db:"study_program_id"`
	SignaturePath      sql.NullString `db:"signature_path"`
	SignaturePathType  sql.NullString `db:"signature_path_type"`
	ShowSignature      bool           `db:"show_signature"`
	EmployeeNo         sql.NullString `db:"employee_no"`
	UpdatedBy          string         `db:"updated_by"`
}
