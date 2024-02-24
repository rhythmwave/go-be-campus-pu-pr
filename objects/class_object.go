package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetClassListRequest struct {
	StudyProgramId         string
	SemesterId             string
	IsActive               *bool
	ClassName              string
	SubjectName            string
	SubjectId              string
	IsMbkm                 bool
	ForOddSemester         *bool
	FollowSemesterIdParity bool
}

type GetClassLecturer struct {
	Id         string
	Name       string
	FrontTitle *string
	BackDegree *string
}

type GetClass struct {
	Id                          string
	Name                        string
	SubjectId                   string
	SubjectCode                 string
	SubjectName                 string
	SubjectIsMandatory          bool
	SubjectSemesterPackage      uint32
	MaximumParticipant          *uint32
	TotalParticipant            uint32
	SubjectTheoryCredit         uint32
	SubjectPracticumCredit      uint32
	SubjectFieldPracticumCredit uint32
	SubjectTotalLessonPlan      uint32
	IsActive                    bool
	UnapprovedStudyPlan         uint32
	TotalMaterial               uint32
	TotalWork                   uint32
	TotalDiscussion             uint32
	TotalExam                   uint32
	TotalEvent                  uint32
	TotalLecturePlan            uint32
	TotalLectureDone            uint32
	TotalGradedParticipant      uint32
	StudyLevelId                string
	ApplicationDeadline         *time.Time
	CurriculumId                string
	CurriculumName              string
	StudyProgramId              string
	StudyProgramName            string
	SemesterId                  string
	SemesterStartYear           uint32
	SchoolYear                  string
	SemesterType                string
	Lecturers                   []GetClassLecturer
}

type ClassListWithPagination struct {
	Pagination common.Pagination
	Data       []GetClass
}

type GetClassParticipantGrade struct {
	ClassGradeComponentId   string
	ClassGradeComponentName string
	InitialGrade            float64
	FinalGrade              float64
}

type GetClassParticipant struct {
	StudentId             string
	StudentNimNumber      int64
	StudentName           string
	StudyProgramId        string
	StudyProgramName      string
	DiktiStudyProgramCode string
	DiktiStudyProgramType string
	StudyLevelShortName   string
	TotalAttendance       uint32
	AttendancePercentage  float64
	TotalSick             uint32
	TotalLeave            uint32
	TotalAwol             uint32
	SubjectName           string
	IsAttend              *bool
	IsSick                *bool
	IsLeave               *bool
	IsAwol                *bool
	GradePoint            float64
	GradeCode             *string
	GradedByAdminId       *string
	GradedByAdminName     *string
	GradedByLecturerId    *string
	GradedByLecturerName  *string
	GradedAt              *time.Time
	SubjectRepetition     uint32
	Grades                []GetClassParticipantGrade
}

type ClassParticipantWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassParticipant
}

type GetClassDetailLecturer struct {
	Id                   string
	Name                 string
	FrontTitle           *string
	BackDegree           *string
	IsGradingResponsible bool
}

type GetClassDetailGradeComponent struct {
	Id         string
	Name       string
	Percentage float64
}

type GetClassDetailStudentGrade struct {
	ClassGradeComponentId   string
	ClassGradeComponentName string
	InitialGrade            float64
	FinalGrade              float64
}

type GetClassDetailStudent struct {
	Id         string
	NimNumber  int64
	Name       string
	GradePoint float64
	GradeCode  *string
	Grades     []GetClassDetailStudentGrade
}

type GetClassDetailGradeType struct {
	Id                  string
	StudyLevelId        string
	StudyLevelShortName string
	Code                string
	GradePoint          float64
	MinimumGrade        float64
	MaximumGrade        float64
	GradeCategory       string
	GradePointCategory  float64
	Label               *string
	EnglishLabel        *string
	StartDate           time.Time
	EndDate             time.Time
}

type GetClassDetail struct {
	Id                    string
	Name                  string
	StudyProgramId        string
	StudyProgramName      string
	DiktiStudyProgramType string
	StudyLevelId          string
	StudyLevelShortName   string
	CurriculumId          string
	CurriculumName        string
	CurriculumYear        string
	SemesterId            string
	SemesterStartYear     uint32
	SchoolYear            string
	SemesterType          string
	GradingStartDate      *time.Time
	GradingEndDate        *time.Time
	SubjectId             string
	SubjectCode           string
	SubjectName           string
	Scope                 *string
	IsOnline              *bool
	IsOffline             *bool
	MinimumParticipant    *uint32
	MaximumParticipant    *uint32
	TotalParticipant      uint32
	Remarks               string
	IsActive              bool
	ApplicationDeadline   *time.Time
	IsGradingResponsible  *bool
	Lecturers             []GetClassDetailLecturer
	GradeComponents       []GetClassDetailGradeComponent
	Students              []GetClassDetailStudent
	GradeTypes            []GetClassDetailGradeType
}

type CreateClassLecturer struct {
	Id                   string
	IsGradingResponsible bool
}

type CreateClass struct {
	SubjectId           string
	SemesterId          string
	Name                string
	Scope               string
	IsOnline            bool
	IsOffline           bool
	MinimumParticipant  uint32
	MaximumParticipant  uint32
	Remarks             string
	ApplicationDeadline time.Time
	Lecturers           []CreateClassLecturer
}

type UpdateClassLecturer struct {
	Id                   string
	IsGradingResponsible bool
}

type UpdateClass struct {
	Id                  string
	SubjectId           string
	Name                string
	Scope               string
	IsOnline            bool
	IsOffline           bool
	MinimumParticipant  uint32
	MaximumParticipant  uint32
	Remarks             string
	ApplicationDeadline time.Time
	Lecturers           []UpdateClassLecturer
}

type UpdateClassMaximumParticipant struct {
	Id                 string
	MaximumParticipant uint32
}
