package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetLectureExamSupervisor struct {
	Id                     string
	Name                   string
	FrontTitle             *string
	BackDegree             *string
	ExamSupervisorRoleId   string
	ExamSupervisorRoleName string
}

type GetLecture struct {
	Id                               string
	LecturePlanDate                  time.Time
	LecturePlanDayOfWeek             uint32
	LecturePlanStartTime             uint32
	LecturePlanEndTime               uint32
	LectureActualDate                *time.Time
	LectureActualDayOfWeek           *uint32
	LectureActualStartTime           *uint32
	LectureActualEndTime             *uint32
	LecturerId                       *string
	LecturerName                     *string
	ForeignLecturerName              *string
	ForeignLecturerSourceInstance    *string
	IsOriginalLecturer               *bool
	IsManualParticipation            *bool
	AutonomousParticipationStartTime *time.Time
	AutonomousParticipationEndTime   *time.Time
	AttendingParticipant             uint32
	UpdatedAt                        *time.Time
	ClassId                          *string
	ClassName                        *string
	RoomId                           string
	RoomName                         *string
	BuildingId                       string
	BuildingName                     string
	IsMidtermExam                    *bool
	IsEndtermExam                    *bool
	IsTheoryExam                     *bool
	IsPracticumExam                  *bool
	IsFieldPracticumExam             *bool
	SubjectCode                      string
	SubjectName                      string
	TotalParticipant                 uint32
	ExamSupervisors                  []GetLectureExamSupervisor
}

type LectureListWithPagination struct {
	Pagination common.Pagination
	Data       []GetLecture
}

type GetLectureDetailStudent struct {
	Id        string
	NimNumber int64
	Name      string
	IsAttend  *bool
	IsSick    *bool
	IsLeave   *bool
	IsAwol    *bool
}

type GetLectureDetail struct {
	Id                     string
	LecturePlanDate        time.Time
	LecturePlanDayOfWeek   uint32
	LecturePlanStartTime   uint32
	LecturePlanEndTime     uint32
	LectureActualDate      *time.Time
	LectureActualDayOfWeek *uint32
	LectureActualStartTime *uint32
	LectureActualEndTime   *uint32
	ClassId                *string
	ClassName              *string
	SubjectId              string
	SubjectName            string
	StudyProgramId         string
	StudyProgramName       string
	SemesterId             string
	SemesterSchoolYear     string
	SemesterType           string
	LectureTheme           *string
	LectureSubject         *string
	Remarks                *string
	Students               []GetLectureDetailStudent
}

type CreateLecturePlanExamSupervisor struct {
	ExamSupervisorId     string
	ExamSupervisorRoleId string
}

type CreateLecturePlan struct {
	LecturePlanDate      time.Time
	LecturePlanStartTime uint32
	LecturePlanEndTime   uint32
	RoomId               string
	LecturerId           string
	IsExam               bool
	IsTheoryExam         bool
	IsPracticumExam      bool
	IsFieldPracticumExam bool
	IsMidtermExam        bool
	IsEndtermExam        bool
	ExamSupervisors      []CreateLecturePlanExamSupervisor
}

type CreateLecture struct {
	ClassId      string
	LecturePlans []CreateLecturePlan
}

type UpdateLectureParticipant struct {
	StudentId string
	IsAttend  bool
	IsSick    bool
	IsLeave   bool
	IsAwol    bool
}

type UpdateLectureExamSupervisor struct {
	ExamSupervisorId     string
	ExamSupervisorRoleId string
}

type UpdateLecture struct {
	Id                             string
	RoomId                         string
	LecturerId                     string
	ForeignLecturerName            string
	ForeignLecturerSourceInstance  string
	LecturePlanDate                time.Time
	LecturePlanStartTime           uint32
	LecturePlanEndTime             uint32
	LectureTheme                   string
	LectureSubject                 string
	Remarks                        string
	IsManualParticipation          bool
	AutonomousParticipationEndTime time.Time
	Participants                   []UpdateLectureParticipant
	ExamSupervisors                []UpdateLectureExamSupervisor
}

type GetLectureParticipation struct {
	Id                     string
	LecturePlanDate        time.Time
	LecturePlanDayOfWeek   uint32
	LecturePlanStartTime   uint32
	LecturePlanEndTime     uint32
	LectureActualDate      *time.Time
	LectureActualDayOfWeek *uint32
	LectureActualStartTime *uint32
	LectureActualEndTime   *uint32
	IsAttend               *bool
	IsSick                 *bool
	IsLeave                *bool
	IsAwol                 *bool
}

type LectureParticipationWithPagination struct {
	Pagination common.Pagination
	Data       []GetLectureParticipation
}

type GetDetailLectureLecture struct {
	LecturePlanDate time.Time
	DayOfWeek       uint32
	StartTime       uint32
	EndTime         uint32
	RoomName        *string
	IsMidtermExam   *bool
	IsEndtermExam   *bool
}

type GetDetailLecture struct {
	StudyProgramName                string
	SubjectCode                     string
	SubjectName                     string
	SemesterPackage                 uint32
	TheoryCredit                    uint32
	PracticumCredit                 uint32
	FieldPracticumCredit            uint32
	SubjectMinimumPassingGradePoint float64
	SubjectIsMandatory              bool
	MaximumParticipant              *uint32
	PrerequisiteSubjects            []string
	Lectures                        []GetDetailLectureLecture
}

type GetLectureCalendarRequest struct {
	Year       uint32
	Month      uint32
	RoomId     string
	LecturerId string
	ClassId    string
}

type GetLectureCalendarLecture struct {
	LecturePlanStartTime          uint32
	LecturePlanEndTime            uint32
	ClassId                       string
	ClassName                     string
	RoomId                        string
	RoomName                      string
	LecturerId                    *string
	LecturerName                  *string
	LecturerFrontTitle            *string
	LecturerBackDegree            *string
	ForeignLecturerName           *string
	ForeignLecturerSourceInstance *string
}

type GetLectureCalendar struct {
	Date     time.Time
	Lectures []GetLectureCalendarLecture
}

type GetLectureHistory struct {
	Id          string
	LectureDate time.Time
	SubjectName string
	AttendTime  time.Time
}

type GetLectureHistoryWithPagination struct {
	Pagination common.Pagination
	Data       []GetLectureHistory
}
