package objects

import "github.com/sccicitb/pupr-backend/objects/common"

type GetStudyLevel struct {
	Id                    string
	Name                  string
	ShortName             string
	KkniQualification     *string
	AcceptanceRequirement *string
	FurtherEducationLevel *string
	ProfessionalStatus    *string
	CourseLanguage        *string
}

type StudyLevelListWithPagination struct {
	Pagination common.Pagination
	Data       []GetStudyLevel
}

type UpdateStudyLevelSkpi struct {
	Id                    string
	KkniQualification     string
	AcceptanceRequirement string
	FurtherEducationLevel string
	ProfessionalStatus    string
	CourseLanguage        string
}
