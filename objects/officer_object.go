package objects

import (
	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetOfficer struct {
	Id                 string
	IdNationalLecturer *string
	Name               string
	Title              *string
	EnglishTitle       *string
	StudyProgramId     *string
	StudyProgramName   *string
	SignaturePath      *string
	SignaturePathType  *string
	SignatureUrl       string
	ShowSignature      bool
	EmployeeNo         *string
}

type OfficerListWithPagination struct {
	Pagination common.Pagination
	Data       []GetOfficer
}

type CreateOfficer struct {
	IdNationalLecturer string
	Name               string
	Title              string
	EnglishTitle       string
	StudyProgramId     string
	SignaturePath      string
	SignaturePathType  string
	ShowSignature      bool
	EmployeeNo         string
}

type UpdateOfficer struct {
	Id                 string
	IdNationalLecturer string
	Name               string
	Title              string
	EnglishTitle       string
	StudyProgramId     string
	SignaturePath      string
	SignaturePathType  string
	ShowSignature      bool
	EmployeeNo         string
}
