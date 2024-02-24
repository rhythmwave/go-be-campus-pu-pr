package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetStudentClassSchedule struct {
	Date      time.Time
	StartTime uint32
	EndTime   uint32
	RoomId    string
	RoomName  *string
}

type GetStudentClass struct {
	Id                          string
	ClassId                     string
	ClassName                   string
	SubjectId                   string
	SubjectCode                 string
	SubjectName                 string
	SubjectTheoryCredit         uint32
	SubjectPracticumCredit      uint32
	SubjectFieldPracticumCredit uint32
	SubjectRepetition           uint32
	SubjectIsMandatory          bool
	TotalAttendance             uint32
	TotalSick                   uint32
	TotalLeave                  uint32
	TotalAwol                   uint32
	GradePoint                  float64
	GradeCode                   *string
	GradedByAdminId             *string
	GradedByAdminName           *string
	GradedByLecturerId          *string
	GradedByLecturerName        *string
	GradedAt                    *string
	AttendancePercentage        float64
	TotalLecture                uint32
	Schedules                   []GetStudentClassSchedule
}

type StudentClassListWithPagination struct {
	Pagination common.Pagination
	Data       []GetStudentClass
}

type TransferStudentClassData struct {
	StudentId          string
	DestinationClassId string
}

type TransferStudentClass struct {
	SourceClassId string
	Data          []TransferStudentClassData
}

type ReshuffleStudentClassStudent struct {
	SourceClassId string
	StudentId     string
}

type ReshuffleStudentClass struct {
	DestinationClassId string
	Students           []ReshuffleStudentClassStudent
}

type MergeStudentClass struct {
	SourceClassIds     []string
	DestinationClassId string
}

type GradeStudentClass struct {
	StudentId             string
	ClassGradeComponentId string
	InitialGrade          float64
}
