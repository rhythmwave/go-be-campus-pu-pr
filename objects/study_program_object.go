package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetStudyProgram struct {
	Id                    string
	Name                  string
	StudyLevelShortName   string
	StudyLevelName        string
	DiktiStudyProgramType string
	DiktiStudyProgramCode string
	Accreditation         *string
	ActiveCurriculumYear  *string
	Degree                *string
	ShortDegree           *string
	EnglishDegree         *string
}

type StudyProgramListWithPagination struct {
	Pagination common.Pagination
	Data       []GetStudyProgram
}

type GetStudyProgramDetail struct {
	Id                            string
	DiktiStudyProgramId           string
	DiktiStudyProgramCode         string
	DiktiStudyProgramName         string
	DiktiStudyProgramType         string
	StudyLevelShortName           string
	StudyLevelName                string
	Name                          string
	EnglishName                   *string
	ShortName                     *string
	EnglishShortName              *string
	AdministrativeUnit            *string
	FacultyId                     string
	FacultyName                   string
	MajorId                       string
	MajorName                     string
	Address                       *string
	PhoneNumber                   *string
	Fax                           *string
	Email                         *string
	Website                       *string
	ContactPerson                 *string
	CuriculumReviewFrequency      string
	CuriculumReviewMethod         string
	EstablishmentDate             time.Time
	IsActive                      bool
	StartSemester                 *string
	OperationalPermitNumber       string
	OperationalPermitDate         time.Time
	OperationalPermitDueDate      time.Time
	HeadLecturerId                *string
	HeadLecturerName              *string
	HeadLecturerMobilePhoneNumber *string
	OperatorName                  *string
	OperatorPhoneNumber           *string
	MinimumGraduationCredit       uint32
	MinimumThesisCredit           uint32
}

type CreateStudyProgram struct {
	DiktiStudyProgramId      string
	Name                     string
	EnglishName              string
	ShortName                string
	EnglishShortName         string
	AdministrativeUnit       string
	MajorId                  string
	Address                  string
	PhoneNumber              string
	Fax                      string
	Email                    string
	Website                  string
	ContactPerson            string
	CuriculumReviewFrequency string
	CuriculumReviewMethod    string
	EstablishmentDate        time.Time
	IsActive                 bool
	StartSemester            string
	OperationalPermitNumber  string
	OperationalPermitDate    time.Time
	OperationalPermitDueDate time.Time
	HeadLecturerId           string
	OperatorName             string
	OperatorPhoneNumber      string
}

type UpdateStudyProgram struct {
	Id                       string
	DiktiStudyProgramId      string
	Name                     string
	EnglishName              string
	ShortName                string
	EnglishShortName         string
	AdministrativeUnit       string
	MajorId                  string
	Address                  string
	PhoneNumber              string
	Fax                      string
	Email                    string
	Website                  string
	ContactPerson            string
	CuriculumReviewFrequency string
	CuriculumReviewMethod    string
	EstablishmentDate        time.Time
	IsActive                 bool
	StartSemester            string
	OperationalPermitNumber  string
	OperationalPermitDate    time.Time
	OperationalPermitDueDate time.Time
	HeadLecturerId           string
	OperatorName             string
	OperatorPhoneNumber      string
	MinimumGraduationCredit  uint32
	MinimumThesisCredit      uint32
}

type UpdateDegreeStudyProgram struct {
	StudyProgramId string
	Degree         string
	ShortDegree    string
	EnglishDegree  string
}
