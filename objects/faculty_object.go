package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetFaculty struct {
	Id        string
	Name      string
	ShortName *string
}

type FacultyListWithPagination struct {
	Pagination common.Pagination
	Data       []GetFaculty
}

type GetFacultyDetail struct {
	Id                        string
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

type CreateFaculty struct {
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

type UpdateFaculty struct {
	Id                        string
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
