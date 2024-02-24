package models

import (
	"time"
)

type CreateYudiciumStudent struct {
	StudentId       string    `db:"student_id"`
	ApplicationDate time.Time `db:"application_date"`
	WithThesis      bool      `db:"with_thesis"`
}

type GetListStudentYudicium struct {
	Id                    string    `db:"id"`
	NimNumber             int64     `db:"nim_number"`
	Name                  string    `db:"name"`
	DiktiStudyProgramCode *string   `db:"dikti_study_program_code"`
	StudyProgramName      *string   `db:"study_program_name"`
	StudyLevelShortName   *string   `db:"study_level_short_name"`
	DiktiStudyProgramType *string   `db:"dikti_study_program_type"`
	TotalCredit           *uint32   `db:"total_credit"`
	Gpa                   *float64  `db:"gpa"`
	Status                *string   `db:"status"`
	ApplicationDate       time.Time `db:"application_date"`
	DoneYudicium          bool      `db:"done_yudicium"`
}
