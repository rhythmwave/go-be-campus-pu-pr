package models

import (
	"database/sql"
	"time"
)

type GetCurriculum struct {
	Id                           string `db:"id"`
	StudyProgramId               string `db:"study_program_id"`
	DiktiStudyProgramCode        string `db:"dikti_study_program_code"`
	StudyProgramName             string `db:"study_program_name"`
	Name                         string `db:"name"`
	Year                         string `db:"year"`
	IdealStudyPeriod             uint32 `db:"ideal_study_period"`
	MaximumStudyPeriod           uint32 `db:"maximum_study_period"`
	IsActive                     bool   `db:"is_active"`
	TotalSubject                 uint32 `db:"total_subject"`
	TotalSubjectWithPrerequisite uint32 `db:"total_subject_with_prerequisite"`
	TotalSubjectWithEquivalence  uint32 `db:"total_subject_with_equivalence"`
}

type GetCurriculumDetail struct {
	Id                    string     `db:"id"`
	StudyProgramId        string     `db:"study_program_id"`
	StudyProgramName      string     `db:"study_program_name"`
	StudyLevelShortName   string     `db:"study_level_short_name"`
	DiktiStudyProgramType string     `db:"dikti_study_program_type"`
	Name                  string     `db:"name"`
	Year                  string     `db:"year"`
	RectorDecisionNumber  *string    `db:"rector_decision_number"`
	RectorDecisionDate    *time.Time `db:"rector_decision_date"`
	AggreeingParty        *string    `db:"aggreeing_party"`
	AggreementDate        *time.Time `db:"aggreement_date"`
	IdealStudyPeriod      uint32     `db:"ideal_study_period"`
	MaximumStudyPeriod    uint32     `db:"maximum_study_period"`
	Remarks               *string    `db:"remarks"`
	IsActive              bool       `db:"is_active"`
	FinalScoreDeterminant string     `db:"final_score_determinant"`
}

type CreateCurriculum struct {
	StudyProgramId        string         `db:"study_program_id"`
	Name                  string         `db:"name"`
	Year                  string         `db:"year"`
	RectorDecisionNumber  sql.NullString `db:"rector_decision_number"`
	RectorDecisionDate    sql.NullTime   `db:"rector_decision_date"`
	AggreeingParty        sql.NullString `db:"aggreeing_party"`
	AggreementDate        sql.NullTime   `db:"aggreement_date"`
	IdealStudyPeriod      uint32         `db:"ideal_study_period"`
	MaximumStudyPeriod    uint32         `db:"maximum_study_period"`
	Remarks               sql.NullString `db:"remarks"`
	IsActive              bool           `db:"is_active"`
	FinalScoreDeterminant string         `db:"final_score_determinant"`
	CreatedBy             string         `db:"created_by"`
}

type UpdateCurriculum struct {
	Id                    string         `db:"id"`
	Name                  string         `db:"name"`
	Year                  string         `db:"year"`
	RectorDecisionNumber  sql.NullString `db:"rector_decision_number"`
	RectorDecisionDate    sql.NullTime   `db:"rector_decision_date"`
	AggreeingParty        sql.NullString `db:"aggreeing_party"`
	AggreementDate        sql.NullTime   `db:"aggreement_date"`
	IdealStudyPeriod      uint32         `db:"ideal_study_period"`
	MaximumStudyPeriod    uint32         `db:"maximum_study_period"`
	Remarks               sql.NullString `db:"remarks"`
	IsActive              bool           `db:"is_active"`
	FinalScoreDeterminant string         `db:"final_score_determinant"`
	UpdatedBy             string         `db:"updated_by"`
}
