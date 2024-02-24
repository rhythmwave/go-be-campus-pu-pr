package objects

import "time"

type GetDetailStudentSubjectSubject struct {
	SemesterType                string
	SemesterStartYear           uint32
	SemesterSchoolYear          string
	SubjectCode                 string
	SubjectName                 string
	SubjectTheoryCredit         uint32
	SubjectPracticumCredit      uint32
	SubjectFieldPracticumCredit uint32
	GradePoint                  float64
	GradeCode                   *string
}

type GetDetailStudentSubject struct {
	Id               string
	Name             string
	NimNumber        int64
	StudyProgramName *string
	TotalCredit      *uint32
	Gpa              *float64
	Subjects         []GetDetailStudentSubjectSubject
}

type GetTranscriptDetailSemesterSubject struct {
	SubjectCode        string
	SubjectName        string
	SubjectEnglishName *string
	TheoryCredit       uint32
	PracticumCredit    uint32
	GradeCode          *string
}

type GetTranscriptDetailSemester struct {
	SemesterPackage uint32
	Subjects        []GetTranscriptDetailSemesterSubject
}

type GetTranscriptDetail struct {
	NimNumber           int64
	Name                string
	BirthRegencyName    *string
	BirthDate           *time.Time
	GraduationDate      *time.Time
	DiplomaNumber       *string
	StudyProgramName    *string
	StudyLevelName      *string
	StudyLevelShortName *string
	TotalCredit         uint32
	Gpa                 *float64
	GraduationPredicate *string
	TheoryCredit        uint32
	PracticumCredit     uint32
	ThesisTitle         *string
	ThesisEnglishTitle  *string
	Semesters           []GetTranscriptDetailSemester
}

type ConvertMbkmGrade struct {
	StudentId             string
	MbkmClassId           string
	DestinationSubjectIds []string
}
