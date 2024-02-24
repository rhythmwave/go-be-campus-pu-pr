package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetClassExam struct {
	Id                     string
	Title                  string
	Abstraction            *string
	FileUrl                string
	LecturerId             string
	LecturerName           string
	LecturerFrontTitle     *string
	LecturerBackDegree     *string
	StartTime              time.Time
	EndTime                time.Time
	TotalSubmission        uint32
	SubmissionFileUrl      string
	SubmissionFilePath     *string
	SubmissionFilePathType *string
	SubmissionPoint        *float64
}

type ClassExamListWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassExam
}

type GetClassExamSubmission struct {
	Id               *string
	StudentId        string
	NimNumber        int64
	Name             string
	StudyProgramName string
	FileUrl          string
	Point            *float64
}

type ClassExamSubmissionWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassExamSubmission
}

type CreateClassExam struct {
	ClassId      string
	Title        string
	Abstraction  string
	FilePath     string
	FilePathType string
	StartTime    time.Time
	EndTime      time.Time
}

type UpdateClassExam struct {
	Id           string
	Title        string
	Abstraction  string
	FilePath     string
	FilePathType string
	StartTime    time.Time
	EndTime      time.Time
}

type GradeClassExamSubmission struct {
	Id    string
	Point float64
}
