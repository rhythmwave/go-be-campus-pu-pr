package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetMajor struct {
	Id          string
	FacultyName string
	Name        string
}

type MajorListWithPagination struct {
	Pagination common.Pagination
	Data       []GetMajor
}

type GetMajorDetail struct {
	Id                        string
	FacultyId                 string
	FacultyName               string
	Name                      string
	ShortName                 *string
	EnglishName               *string
	EnglishShortName          *string
	Address                   string
	PhoneNumber               *string
	Fax                       *string
	Email                     *string
	ContactPerson             *string
	ExperimentBuildingArea    *float64
	LectureHallArea           *float64
	LectureHallCount          *uint32
	LaboratoriumArea          *float64
	LaboratoriumCount         *uint32
	PermanentLecturerRoomArea *float64
	AdministrationRoomArea    *float64
	BookCount                 *uint32
	BookCopyCount             *uint32
}

type CreateMajor struct {
	FacultyId                 string
	Name                      string
	ShortName                 string
	EnglishName               string
	EnglishShortName          string
	Address                   string
	PhoneNumber               string
	Fax                       string
	Email                     string
	ContactPerson             string
	ExperimentBuildingArea    float64
	LectureHallArea           float64
	LectureHallCount          uint32
	LaboratoriumArea          float64
	LaboratoriumCount         uint32
	PermanentLecturerRoomArea float64
	AdministrationRoomArea    float64
	BookCount                 uint32
	BookCopyCount             uint32
}

type UpdateMajor struct {
	Id                        string
	FacultyId                 string
	Name                      string
	ShortName                 string
	EnglishName               string
	EnglishShortName          string
	Address                   string
	PhoneNumber               string
	Fax                       string
	Email                     string
	ContactPerson             string
	ExperimentBuildingArea    float64
	LectureHallArea           float64
	LectureHallCount          uint32
	LaboratoriumArea          float64
	LaboratoriumCount         uint32
	PermanentLecturerRoomArea float64
	AdministrationRoomArea    float64
	BookCount                 uint32
	BookCopyCount             uint32
}
