package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type BulkCreateStudyPlan struct {
	SemesterId string
	StudentIds []string
	ClassIds   []string
	IsThesis   bool
}

type GetListStudyPlan struct {
	SemesterId            string
	SemesterStartYear     uint32
	SchoolYear            string
	SemesterType          string
	TotalMandatoryCredit  uint32
	TotalOptionalCredit   uint32
	GradePoint            float64
	StudentId             string
	StudentNimNumber      int64
	StudentName           string
	StudyProgramId        string
	StudyProgramName      string
	DiktiStudyProgramCode string
	DiktiStudyProgramType string
	StudyLevelShortName   string
	IsThesis              bool
}

type GetStudyPlanWithPagination struct {
	Pagination common.Pagination
	Data       []GetListStudyPlan
}

type GetStudentStudyPlanDetailClassSchedule struct {
	Date      time.Time
	StartTime uint32
	EndTime   uint32
	RoomId    string
	RoomName  *string
}

type GetStudentStudyPlanDetailClass struct {
	Id                          string
	Name                        string
	SubjectId                   string
	SubjectName                 string
	SubjectCode                 string
	SubjectTheoryCredit         uint32
	SubjectPracticumCredit      uint32
	SubjectFieldPracticumCredit uint32
	SubjectIsMandatory          bool
	TotalLectureDone            uint32
	TotalAttendance             uint32
	ActiveLectureId             *string
	ActiveLectureHasAttend      *bool
	GradePoint                  float64
	GradeCode                   *string
	Schedules                   []GetStudentStudyPlanDetailClassSchedule
}

type GetStudentStudyPlanDetail struct {
	StudyPlanInputStartDate            time.Time
	StudyPlanInputEndDate              time.Time
	Id                                 string
	IsSubmitted                        bool
	IsApproved                         bool
	StudentId                          string
	StudentNimNumber                   int64
	StudentName                        string
	StudyProgramId                     *string
	StudyProgramName                   *string
	SemesterId                         string
	SemesterSchoolYear                 string
	SemesterType                       string
	MaximumCredit                      uint32
	AcademicGuidanceLecturerId         *string
	AcademicGuidanceLecturerName       *string
	AcademicGuidanceLecturerFrontTitle *string
	AcademicGuidanceLecturerBackDegree *string
	TotalMandatoryCredit               uint32
	TotalOptionalCredit                uint32
	GradePoint                         float64
	Gpa                                *float64
	IsThesis                           bool
	Classes                            []GetStudentStudyPlanDetailClass
}
