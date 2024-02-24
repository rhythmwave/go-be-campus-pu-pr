package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetSemesterCurriculum struct {
	StudyProgramId   string
	StudyProgramName string
	CurriculumId     string
	CurriculumName   string
}

type GetSemester struct {
	Id                string
	SemesterStartYear uint32
	SchoolYear        string
	SemesterType      string
	IsActive          bool
	StartDate         time.Time
	EndDate           time.Time
	MidtermStartDate  *time.Time
	MidtermEndDate    *time.Time
	EndtermStartDate  *time.Time
	EndtermEndDate    *time.Time
	Curriculums       []GetSemesterCurriculum
}

type SemesterListWithPagination struct {
	Pagination common.Pagination
	Data       []GetSemester
}

type GetSemesterDetailCurriculum struct {
	StudyProgramId   string
	StudyProgramName string
	CurriculumId     string
	CurriculumName   string
}

type GetSemesterDetail struct {
	Id                         string
	SemesterStartYear          uint32
	SchoolYear                 string
	SemesterType               string
	IsActive                   bool
	StartDate                  time.Time
	EndDate                    time.Time
	MidtermStartDate           *time.Time
	MidtermEndDate             *time.Time
	EndtermStartDate           *time.Time
	EndtermEndDate             *time.Time
	StudyPlanInputStartDate    time.Time
	StudyPlanInputEndDate      time.Time
	StudyPlanApprovalStartDate time.Time
	StudyPlanApprovalEndDate   time.Time
	GradingStartDate           *time.Time
	GradingEndDate             *time.Time
	ReferenceSemesterId        *string
	ReferenceSemesterStartYear *uint32
	ReferenceSchoolYear        *string
	ReferenceSemesterType      *string
	CheckMinimumGpa            bool
	CheckPassedCredit          bool
	DefaultCredit              uint32
	Curriculums                []GetSemesterDetailCurriculum
}

type CreateSemesterCurriculum struct {
	CurriculumId string
}

type CreateSemester struct {
	SemesterStartYear          uint32
	SemesterType               string
	StartDate                  time.Time
	EndDate                    time.Time
	MidtermStartDate           time.Time
	MidtermEndDate             time.Time
	EndtermStartDate           time.Time
	EndtermEndDate             time.Time
	StudyPlanInputStartDate    time.Time
	StudyPlanInputEndDate      time.Time
	StudyPlanApprovalStartDate time.Time
	StudyPlanApprovalEndDate   time.Time
	GradingStartDate           time.Time
	GradingEndDate             time.Time
	ReferenceSemesterId        string
	CheckMinimumGpa            bool
	CheckPassedCredit          bool
	DefaultCredit              uint32
	Curriculums                []CreateSemesterCurriculum
}

type UpdateSemesterCurriculum struct {
	CurriculumId string
}
type UpdateSemester struct {
	Id                         string
	SemesterStartYear          uint32
	SemesterType               string
	StartDate                  time.Time
	EndDate                    time.Time
	MidtermStartDate           time.Time
	MidtermEndDate             time.Time
	EndtermStartDate           time.Time
	EndtermEndDate             time.Time
	StudyPlanInputStartDate    time.Time
	StudyPlanInputEndDate      time.Time
	StudyPlanApprovalStartDate time.Time
	StudyPlanApprovalEndDate   time.Time
	GradingStartDate           time.Time
	GradingEndDate             time.Time
	ReferenceSemesterId        string
	CheckMinimumGpa            bool
	CheckPassedCredit          bool
	DefaultCredit              uint32
	Curriculums                []UpdateSemesterCurriculum
}
