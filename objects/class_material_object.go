package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetClassMaterial struct {
	Id                 string
	Title              string
	Abstraction        *string
	FilePath           *string
	FilePathType       *string
	FileUrl            string
	LecturerId         string
	LecturerName       string
	LecturerFrontTitle *string
	LecturerBackDegree *string
	IsActive           bool
	CreatedAt          string
}

type ClassMaterialListWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassMaterial
}

type CreateClassMaterial struct {
	ClassId      string
	Title        string
	Abstraction  string
	FilePath     string
	FilePathType string
	IsActive     bool
}

type UpdateClassMaterial struct {
	Id           string
	Title        string
	Abstraction  string
	FilePath     string
	FilePathType string
	IsActive     bool
}
