package models

import "database/sql"

type GetStudyLevel struct {
	Id                    string  `db:"id"`
	Name                  string  `db:"name"`
	ShortName             string  `db:"short_name"`
	KkniQualification     *string `db:"kkni_qualification"`
	AcceptanceRequirement *string `db:"acceptance_requirement"`
	FurtherEducationLevel *string `db:"further_education_level"`
	ProfessionalStatus    *string `db:"professional_status"`
	CourseLanguage        *string `db:"course_language"`
}

type UpdateStudyLevelSkpi struct {
	Id                    string         `db:"id"`
	KkniQualification     sql.NullString `db:"kkni_qualification"`
	AcceptanceRequirement sql.NullString `db:"acceptance_requirement"`
	FurtherEducationLevel sql.NullString `db:"further_education_level"`
	ProfessionalStatus    sql.NullString `db:"professional_status"`
	CourseLanguage        sql.NullString `db:"course_language"`
	UpdatedBy             string         `db:"updated_by"`
}
