package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetStudentLeaveRequest struct {
	Id                         string
	NimNumber                  int64
	Name                       string
	DiktiStudyProgramCode      string
	StudyProgramName           string
	StudyLevelShortName        string
	DiktiStudyProgramType      string
	StartDate                  time.Time
	TotalLeaveDurationSemester uint32
	PermitNumber               *string
	Purpose                    string
	Remarks                    string
	IsApproved                 *bool
	SemesterType               string
	SemesterSchoolYear         string
}

type StudentLeaveRequestListWithPagination struct {
	Pagination common.Pagination
	Data       []GetStudentLeaveRequest
}

type GetStudentLeave struct {
	Id                    string
	NimNumber             int64
	Name                  string
	DiktiStudyProgramCode string
	StudyProgramName      string
	StudyLevelShortName   string
	DiktiStudyProgramType string
	SemesterSchoolYear    string
	SemesterType          string
	PermitNumber          *string
	Purpose               string
	Remarks               string
}

type StudentLeaveListWithPagination struct {
	Pagination common.Pagination
	Data       []GetStudentLeave
}

type CreateStudentLeave struct {
	StudentId                  string
	TotalLeaveDurationSemester uint32
	StartDate                  time.Time
	PermitNumber               string
	Purpose                    string
	Remarks                    string
}

type UpdateStudentLeave struct {
	Id           string
	PermitNumber string
	Purpose      string
	Remarks      string
}
