package models

import "time"

type GetLecturerResignation struct {
	Id                    string    `db:"id"`
	Name                  string    `db:"name"`
	IdNationalLecturer    string    `db:"id_national_lecturer"`
	FrontTitle            *string   `db:"front_title"`
	BackDegree            *string   `db:"back_degree"`
	SemesterStartYear     uint32    `db:"semester_start_year"`
	SemesterType          string    `db:"semester_type"`
	DiktiStudyProgramCode *string   `db:"dikti_study_program_code"`
	StudyProgramName      *string   `db:"study_program_name"`
	StudyLevelShortName   *string   `db:"study_level_short_name"`
	DiktiStudyProgramType *string   `db:"dikti_study_program_type"`
	ResignDate            time.Time `db:"resign_date"`
	ResignationNumber     string    `db:"resignation_number"`
	Purpose               string    `db:"purpose"`
	Remarks               string    `db:"remarks"`
}

type CreateLecturerResignation struct {
	LecturerId        string    `db:"lecturer_id"`
	SemesterId        string    `db:"semester_id"`
	ResignDate        time.Time `db:"resign_date"`
	ResignationNumber string    `db:"resignation_number"`
	Purpose           string    `db:"purpose"`
	Remarks           string    `db:"remarks"`
	CreatedBy         string    `db:"created_by"`
}
