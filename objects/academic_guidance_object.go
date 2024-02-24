package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetAcademicGuidanceStudent struct {
	Id                      string
	NimNumber               int64
	StudentForce            *string
	Name                    string
	Status                  *string
	StudyPlanFormIsApproved *bool
}

type AcademicGuidanceStudentListWithPagination struct {
	Pagination common.Pagination
	Data       []GetAcademicGuidanceStudent
}

type GetAcademicGuidanceDetail struct {
	Id                 string
	SemesterId         string
	LecturerId         string
	LecturerName       string
	LecturerFrontTitle *string
	LecturerBackDegree *string
	DecisionNumber     *string
	DecisionDate       *time.Time
	TotalStudent       uint32
}

type UpsertAcademicGuidance struct {
	SemesterId string
	LecturerId string
	StudentIds []string
}

type UpsertDecisionAcademicGuidance struct {
	LecturerId     string
	SemesterId     string
	DecisionNumber string
	DecisionDate   time.Time
}

type GetAcademicGuidanceSessionFile struct {
	Id           string
	Title        string
	FileUrl      string
	FilePath     string
	FilePathType string
}

type GetAcademicGuidanceSessionStudent struct {
	Id        string
	Name      string
	NimNumber int64
}

type GetAcademicGuidanceSession struct {
	Id                 string
	AcademicGuidanceId string
	Subject            string
	SessionDate        time.Time
	Summary            string
	Files              []GetAcademicGuidanceSessionFile
	Students           []GetAcademicGuidanceSessionStudent
}

type CreateAcademicGuidanceSessionFile struct {
	Title        string
	FilePath     string
	FilePathType string
}

type CreateAcademicGuidanceSession struct {
	AcademicGuidanceId string
	SemesterId         string
	Subject            string
	SessionDate        time.Time
	Summary            string
	Files              []CreateAcademicGuidanceSessionFile
	StudentIds         []string
}

type UpdateAcademicGuidanceSessionFile struct {
	Title        string
	FilePath     string
	FilePathType string
}

type UpdateAcademicGuidanceSession struct {
	Id          string
	Subject     string
	SessionDate time.Time
	Summary     string
	Files       []UpdateAcademicGuidanceSessionFile
	StudentIds  []string
}
