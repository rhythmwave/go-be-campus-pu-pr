package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetLecturerResignation struct {
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
	ResignDate            time.Time
	ResignationNumber     string
	Purpose               string
	Remarks               string
}

type LecturerResignationListWithPagination struct {
	Pagination common.Pagination
	Data       []GetLecturerResignation
}

type CreateLecturerResignation struct {
	LecturerId        string
	SemesterId        string
	ResignDate        time.Time
	ResignationNumber string
	Purpose           string
	Remarks           string
}
