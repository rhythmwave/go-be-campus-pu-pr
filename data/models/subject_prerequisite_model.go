package models

import "database/sql"

type GetSubjectPrerequisite struct {
	Id                      string   `db:"id"`
	SubjectId               string   `db:"subject_id"`
	PrerequisiteSubjectId   string   `db:"prerequisite_subject_id"`
	PrerequisiteSubjectCode string   `db:"prerequisite_subject_code"`
	PrerequisiteSubjectName string   `db:"prerequisite_subject_name"`
	PrerequisiteType        string   `db:"prerequisite_type"`
	MinimumGradePoint       *float64 `db:"minimum_grade_point"`
}

type CreateSubjectPrerequisite struct {
	SubjectId             string          `db:"subject_id"`
	PrerequisiteSubjectId string          `db:"prerequisite_subject_id"`
	PrerequisiteType      string          `db:"prerequisite_type"`
	MinimumGradePoint     sql.NullFloat64 `db:"minimum_grade_point"`
	CreatedBy             string          `db:"created_by"`
}
