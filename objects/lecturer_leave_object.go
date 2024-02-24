package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetLecturerLeave struct {
	Id                    string
	Name                  string
	IdNationalLecturer    string
	FrontTitle            *string
	BackDegree            *string
	SemesterSchoolYear    string
	SemesterType          string
	DiktiStudyProgramCode *string
	StudyProgramName      *string
	StudyLevelShortName   *string
	DiktiStudyProgramType *string
	StartDate             time.Time
	EndDate               *time.Time
	PermitNumber          string
	Purpose               string
	Remarks               string
	FileUrl               string
	FilePath              *string
	FilePathType          *string
}

type LecturerLeaveListWithPagination struct {
	Pagination common.Pagination
	Data       []GetLecturerLeave
}

type CreateLecturerLeave struct {
	LecturerId   string
	SemesterId   string
	StartDate    time.Time
	PermitNumber string
	Purpose      string
	Remarks      string
	FilePath     string
	FilePathType string
}

type UpdateLecturerLeave struct {
	Id           string
	PermitNumber string
	Purpose      string
	Remarks      string
	FilePath     string
	FilePathType string
}
