package objects

import (
	"time"

	"github.com/sccicitb/pupr-backend/objects/common"
)

type GetStudentActivity struct {
	Id                 string
	StudyProgramId     string
	StudyProgramName   string
	SemesterId         string
	SemesterSchoolYear string
	SemesterType       string
	ActivityType       string
	Title              string
}

type StudentActivityListWithPagination struct {
	Pagination common.Pagination
	Data       []GetStudentActivity
}

type GetStudentActivityDetailParticipant struct {
	StudentId        string
	NimNumber        int64
	Name             string
	StudyProgramId   *string
	StudyProgramName *string
	Role             string
}

type GetStudentActivityDetailLecturer struct {
	LecturerId         string
	IdNationalLecturer string
	Name               string
	FrontTitle         *string
	BackDegree         *string
	ActivityCategory   string
	Sort               uint32
}

type GetStudentActivityDetail struct {
	Id                 string
	StudyProgramId     string
	StudyProgramName   string
	SemesterId         string
	SemesterSchoolYear string
	SemesterType       string
	ActivityType       string
	Title              string
	Location           *string
	DecisionNumber     *string
	DecisionDate       *time.Time
	IsGroupActivity    bool
	Remarks            *string
	Participants       []GetStudentActivityDetailParticipant
	Mentors            []GetStudentActivityDetailLecturer
	Examiners          []GetStudentActivityDetailLecturer
}

type CreateStudentActivityParticipant struct {
	StudentId string
	Role      string
}

type CreateStudentActivityLecturer struct {
	LecturerId       string
	ActivityCategory string
	Sort             uint32
}

type CreateStudentActivity struct {
	StudyProgramId  string
	SemesterId      string
	ActivityType    string
	Title           string
	Location        string
	DecisionNumber  string
	DecisionDate    string
	IsGroupActivity bool
	Remarks         string
	IsMbkm          bool
	Participants    []CreateStudentActivityParticipant
	Mentors         []CreateStudentActivityLecturer
	Examiners       []CreateStudentActivityLecturer
}

type UpdateStudentActivityParticipant struct {
	StudentId string
	Role      string
}

type UpdateStudentActivityLecturer struct {
	LecturerId       string
	ActivityCategory string
	Sort             uint32
}

type UpdateStudentActivity struct {
	Id              string
	StudyProgramId  string
	SemesterId      string
	ActivityType    string
	Title           string
	Location        string
	DecisionNumber  string
	DecisionDate    string
	IsGroupActivity bool
	Remarks         string
	IsMbkm          bool
	Participants    []UpdateStudentActivityParticipant
	Mentors         []UpdateStudentActivityLecturer
	Examiners       []UpdateStudentActivityLecturer
}
