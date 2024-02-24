package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetLecturerMutation struct {
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
	MutationDate          time.Time
	DecisionNumber        string
	Destination           string
}

type LecturerMutationListWithPagination struct {
	Pagination common.Pagination
	Data       []GetLecturerMutation
}

type CreateLecturerMutation struct {
	LecturerId     string
	SemesterId     string
	MutationDate   time.Time
	DecisionNumber string
	Destination    string
}
