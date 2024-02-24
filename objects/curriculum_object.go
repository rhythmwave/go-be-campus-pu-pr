package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetCurriculum struct {
	Id                           string
	StudyProgramId               string
	StudyProgramName             string
	DiktiStudyProgramCode        string
	Name                         string
	Year                         string
	IdealStudyPeriod             uint32
	MaximumStudyPeriod           uint32
	IsActive                     bool
	TotalSubject                 uint32
	TotalSubjectWithPrerequisite uint32
	TotalSubjectWithEquivalence  uint32
}

type CurriculumListWithPagination struct {
	Pagination common.Pagination
	Data       []GetCurriculum
}

type GetCurriculumDetail struct {
	Id                    string
	StudyProgramId        string
	StudyProgramName      string
	StudyLevelShortName   string
	DiktiStudyProgramType string
	Name                  string
	Year                  string
	RectorDecisionNumber  *string
	RectorDecisionDate    *time.Time
	AggreeingParty        *string
	AggreementDate        *time.Time
	IdealStudyPeriod      uint32
	MaximumStudyPeriod    uint32
	Remarks               *string
	IsActive              bool
	FinalScoreDeterminant string
}

type CreateCurriculum struct {
	StudyProgramId        string
	Name                  string
	Year                  string
	RectorDecisionNumber  string
	RectorDecisionDate    time.Time
	AggreeingParty        string
	AggreementDate        time.Time
	IdealStudyPeriod      uint32
	MaximumStudyPeriod    uint32
	Remarks               string
	IsActive              bool
	FinalScoreDeterminant string
}

type UpdateCurriculum struct {
	Id                    string
	Name                  string
	Year                  string
	RectorDecisionNumber  string
	RectorDecisionDate    time.Time
	AggreeingParty        string
	AggreementDate        time.Time
	IdealStudyPeriod      uint32
	MaximumStudyPeriod    uint32
	Remarks               string
	IsActive              bool
	FinalScoreDeterminant string
}
