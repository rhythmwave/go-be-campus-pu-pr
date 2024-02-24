package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetSubjectRequest struct {
	CurriculumIds            []string
	PrerequisiteOfSubjectId  string
	EquivalentToCurriculumId string
	SemesterPackage          uint32
	ClassSemesterId          string
	StudyProgramId           string
	IsThesis                 *bool
	IsMbkm                   bool
}

type GetSubjectClass struct {
	Id   string
	Name string
}

type GetSubject struct {
	Id                            string
	StudyProgramId                string
	StudyProgramName              string
	CurriculumId                  string
	CurriculumName                string
	Code                          string
	Name                          string
	IsMandatory                   bool
	SemesterPackage               uint32
	TheoryCredit                  uint32
	PracticumCredit               uint32
	FieldPracticumCredit          uint32
	SubjectPrerequisiteId         *string
	PrerequisiteType              *string
	PrerequisiteMinimumGradePoint *float64
	EquivalentStudyProgramId      *string
	EquivalentStudyProgramName    *string
	EquivalentCurriculumId        *string
	EquivalentCurriculumName      *string
	EquivalentSubjectId           *string
	EquivalentSubjectCode         *string
	EquivalentSubjectName         *string
	SubjectCategoryId             string
	SubjectCategoryName           string
	IsThesis                      *bool
	IsMbkm                        bool
	TotalLessonPlan               uint32
	Classes                       []GetSubjectClass
}

type SubjectListWithPagination struct {
	Pagination common.Pagination
	Data       []GetSubject
}

type GetSubjectDetailPrerequisite struct {
	Id                string
	Code              string
	Name              string
	PrerequisiteType  string
	MinimumGradePoint *float64
}

type GetSubjectDetail struct {
	Id                           string
	StudyProgramId               string
	StudyProgramName             string
	CurriculumId                 string
	CurriculumName               string
	Code                         string
	Name                         string
	ShortName                    *string
	EnglishName                  *string
	EnglishShortName             *string
	IsMandatory                  bool
	Trait                        string
	Type                         *string
	SubjectCategoryId            string
	SubjectCategoryName          string
	CurriculumType               string
	TheoryCredit                 uint32
	PracticumCredit              uint32
	FieldPracticumCredit         uint32
	SemesterPackage              uint32
	RepeatCourseLimit            uint32
	IsActive                     bool
	HasLectureUnit               bool
	HasTeachingMaterial          bool
	HasLectureSummary            bool
	SupportingLecturerId         *string
	SupportingLecturerName       *string
	StartDate                    *time.Time
	EndDate                      *time.Time
	MinimumPassingGradePoint     float64
	MinimumMandatoryCreditTaken  *uint32
	MinimumOptionalCreditTaken   *uint32
	MinimumTotalCreditTaken      *uint32
	MinimumMandatoryCreditPassed *uint32
	MinimumOptionalCreditPassed  *uint32
	MinimumTotalCreditPassed     *uint32
	MinimumGpa                   *float64
	Abstraction                  *string
	SyllabusPath                 *string
	SyllabusPathType             *string
	SyllabusUrl                  string
	IsThesis                     *bool
	IsMbkm                       bool
	SubjectPrerequisites         []GetSubjectDetailPrerequisite
}

type CreateSubject struct {
	CurriculumId                 string
	Code                         string
	Name                         string
	ShortName                    string
	EnglishName                  string
	EnglishShortName             string
	IsMandatory                  bool
	Trait                        string
	Type                         string
	SubjectCategoryId            string
	CurriculumType               string
	TheoryCredit                 uint32
	PracticumCredit              uint32
	FieldPracticumCredit         uint32
	SemesterPackage              uint32
	RepeatCourseLimit            uint32
	IsActive                     bool
	HasLectureUnit               bool
	HasTeachingMaterial          bool
	HasLectureSummary            bool
	SupportingLecturerId         string
	StartDate                    time.Time
	EndDate                      time.Time
	MinimumPassingGradePoint     float64
	MinimumMandatoryCreditTaken  uint32
	MinimumOptionalCreditTaken   uint32
	MinimumTotalCreditTaken      uint32
	MinimumMandatoryCreditPassed uint32
	MinimumOptionalCreditPassed  uint32
	MinimumTotalCreditPassed     uint32
	MinimumGpa                   float64
	Abstraction                  string
	SyllabusPath                 string
	SyllabusPathType             string
	IsThesis                     bool
	IsMbkm                       bool
}

type UpdateSubject struct {
	Id                           string
	CurriculumId                 string
	Code                         string
	Name                         string
	ShortName                    string
	EnglishName                  string
	EnglishShortName             string
	IsMandatory                  bool
	Trait                        string
	Type                         string
	SubjectCategoryId            string
	CurriculumType               string
	TheoryCredit                 uint32
	PracticumCredit              uint32
	FieldPracticumCredit         uint32
	SemesterPackage              uint32
	RepeatCourseLimit            uint32
	IsActive                     bool
	HasLectureUnit               bool
	HasTeachingMaterial          bool
	HasLectureSummary            bool
	SupportingLecturerId         string
	StartDate                    time.Time
	EndDate                      time.Time
	MinimumPassingGradePoint     float64
	MinimumMandatoryCreditTaken  uint32
	MinimumOptionalCreditTaken   uint32
	MinimumTotalCreditTaken      uint32
	MinimumMandatoryCreditPassed uint32
	MinimumOptionalCreditPassed  uint32
	MinimumTotalCreditPassed     uint32
	MinimumGpa                   float64
	Abstraction                  string
	SyllabusPath                 string
	SyllabusPathType             string
	IsThesis                     bool
	IsMbkm                       bool
}

type SetPrerequisiteSubject struct {
	Id                string
	PrerequisiteType  string
	MinimumGradePoint float64
}

type SetEquivalentSubject struct {
	SubjectId           string
	EquivalentSubjectId string
	IsViceVersa         bool
}
