package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetListThesis struct {
	Id                        string
	Topic                     string
	Title                     string
	Status                    string
	StudentId                 string
	StudentName               string
	StudentNimNumber          int64
	StudentStatus             *string
	StudyProgramId            string
	StudyProgramName          string
	DiktiStudyProgramCode     string
	DiktiStudyProgramType     string
	StudyLevelShortName       string
	StudentHasThesisStudyPlan bool
	StartSemesterId           string
	StartSemesterType         string
	StartSemesterSchoolYear   string
}

type GetListThesisWithPagination struct {
	Pagination common.Pagination
	Data       []GetListThesis
}

type GetDetailThesisFile struct {
	Id              string
	FileUrl         string
	FilePath        string
	FilePathType    string
	FileDescription *string
}

type GetDetailThesisSupervisor struct {
	Id                       string
	LecturerId               string
	LecturerName             string
	LecturerFrontTitle       *string
	LecturerBackDegree       *string
	ThesisSupervisorRoleId   string
	ThesisSupervisorRoleName string
	ThesisSupervisorRoleSort uint32
}

type GetDetailThesis struct {
	Id                        string
	StudyProgramId            *string
	StudentId                 string
	StudentName               string
	StudentNimNumber          int64
	StartSemesterId           string
	StartSemesterType         string
	StartSemesterSchoolYear   string
	FinishSemesterId          *string
	FinishSemesterType        *string
	FinishSemesterSchoolYear  string
	Topic                     string
	Title                     string
	EnglishTitle              *string
	StartDate                 time.Time
	FinishDate                *time.Time
	Remarks                   *string
	IsJointThesis             bool
	Status                    string
	ProposalSeminarDate       *time.Time
	ProposalCertificateNumber *string
	ProposalCertificateDate   *time.Time
	ThesisDefenseCount        uint32
	GradePoint                float64
	GradeCode                 *string
	Files                     []GetDetailThesisFile
	ThesisSupervisors         []GetDetailThesisSupervisor
}

type CreateThesisSupervisor struct {
	LecturerId             string
	ThesisSupervisorRoleId string
}

type CreateThesis struct {
	StudentId                 string
	Topic                     string
	Status                    string
	Title                     string
	EnglishTitle              string
	StartSemesterId           string
	StartDate                 time.Time
	Remarks                   string
	IsJointThesis             bool
	FilePath                  string
	FilePathType              string
	FileDescription           string
	ProposalSeminarDate       time.Time
	ProposalCertificateNumber string
	ProposalCertificateDate   time.Time
	ThesisSupervisors         []CreateThesisSupervisor
}

type UpdateThesisFile struct {
	FilePath        string
	FilePathType    string
	FileDescription string
}

type UpdateThesisSupervisor struct {
	LecturerId             string
	ThesisSupervisorRoleId string
}

type UpdateThesis struct {
	Id                        string
	StudentId                 string
	Topic                     string
	Status                    string
	Title                     string
	EnglishTitle              string
	StartSemesterId           string
	StartDate                 time.Time
	Remarks                   string
	IsJointThesis             bool
	Files                     []UpdateThesisFile
	ProposalSeminarDate       time.Time
	ProposalCertificateNumber string
	ProposalCertificateDate   time.Time
	ThesisSupervisors         []UpdateThesisSupervisor
}

type GetListThesisDefenseRequestExaminer struct {
	Id                     string
	LecturerId             string
	LecturerName           string
	LecturerFrontTitle     *string
	LecturerBackDegree     *string
	ThesisExaminerRoleId   string
	ThesisExaminerRoleName string
}

type GetListThesisDefenseRequest struct {
	Id                           string
	StudentId                    string
	StudentName                  string
	StudentNimNumber             int64
	StudentStatus                string
	StudyProgramId               string
	StudyProgramName             string
	DiktiStudyProgramCode        string
	DiktiStudyProgramType        string
	StudyLevelId                 string
	StudyLevelShortName          string
	ThesisId                     string
	ThesisTitle                  string
	ThesisStatus                 string
	ThesisDefenseCount           uint32
	ThesisDefenseId              *string
	ThesisDefensePlanDate        *time.Time
	ThesisDefensePlanStartTime   *uint32
	ThesisDefensePlanEndTime     *uint32
	ThesisDefenseActualDate      *time.Time
	ThesisDefenseActualStartTime *uint32
	ThesisDefenseActualEndTime   *uint32
	ThesisDefenseRoomId          *string
	ThesisDefenseRoomName        *string
	ThesisDefenseRevision        *string
	ThesisGradeCode              *string
	ThesisDefenseIsPassed        *bool
	CreatedAt                    time.Time
	Examiners                    []GetListThesisDefenseRequestExaminer
}

type GetListThesisDefenseRequestWithPagination struct {
	Pagination common.Pagination
	Data       []GetListThesisDefenseRequest
}

type CreateThesisDefenseExaminer struct {
	LecturerId           string
	ThesisExaminerRoleId string
}

type CreateThesisDefense struct {
	ThesisId      string
	PlanDate      time.Time
	PlanStartTime uint32
	PlanEndTime   uint32
	RoomId        string
	Examiners     []CreateThesisDefenseExaminer
}

type UpdateThesisDefenseExaminer struct {
	LecturerId           string
	ThesisExaminerRoleId string
}

type UpdateThesisDefense struct {
	Id              string
	PlanDate        time.Time
	PlanStartTime   uint32
	PlanEndTime     uint32
	RoomId          string
	ActualDate      time.Time
	ActualStartTime uint32
	ActualEndTime   uint32
	IsPassed        bool
	Revision        string
	GradeCode       string
	Examiners       []UpdateThesisDefenseExaminer
}

type GetThesisSupervisorLogThesisSupervisorRole struct {
	Id    string
	Name  string
	Total uint32
}

type GetThesisSupervisorLog struct {
	Id                     string
	IdNationalLecturer     string
	Name                   string
	TotalSupervisedThesis  uint32
	ActiveSupervisedThesis uint32
	ThesisSupervisorRoles  []GetThesisSupervisorLogThesisSupervisorRole
}

type GetThesisSupervisorLogWithPagination struct {
	Pagination common.Pagination
	Data       []GetThesisSupervisorLog
}
