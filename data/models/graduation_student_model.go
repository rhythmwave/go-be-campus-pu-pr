package models

import (
	"time"
)

type CreateGraduationStudent struct {
	StudentId           string    `db:"student_id"`
	ApplicationDate     time.Time `db:"application_date"`
	GraduationSessionId string    `db:"graduation_session_id"`
}

type GetListStudentGraduation struct {
	Id                    string    `db:"id"`
	NimNumber             int64     `db:"nim_number"`
	Name                  string    `db:"name"`
	DiktiStudyProgramCode *string   `db:"dikti_study_program_code"`
	StudyProgramName      *string   `db:"study_program_name"`
	StudyLevelShortName   *string   `db:"study_level_short_name"`
	DiktiStudyProgramType *string   `db:"dikti_study_program_type"`
	ApplicationDate       time.Time `db:"application_date"`
}
