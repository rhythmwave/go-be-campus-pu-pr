package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetClassWork struct {
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

type ClassWorkListWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassWork
}

type GetClassWorkSubmission struct {
	Id               *string
	StudentId        string
	NimNumber        int64
	Name             string
	StudyProgramName string
	FileUrl          string
	Point            *float64
}

type ClassWorkSubmissionWithPagination struct {
	Pagination common.Pagination
	Data       []GetClassWorkSubmission
}

type CreateClassWork struct {
	ClassId      string
	Title        string
	Abstraction  string
	FilePath     string
	FilePathType string
	StartTime    time.Time
	EndTime      time.Time
}

type UpdateClassWork struct {
	Id           string
	Title        string
	Abstraction  string
	FilePath     string
	FilePathType string
	StartTime    time.Time
	EndTime      time.Time
}

type GradeClassWorkSubmission struct {
	Id    string
	Point float64
}
